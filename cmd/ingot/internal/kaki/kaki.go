package kaki

import (
	"bytes"
	"context"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/decode"
	"log"
	"os"
)

const kakiVersion = 578

// kakiClient implements the protocol interface and holds all needed information for packet decoding / encoding
type kakiClient struct {
	log          *log.Logger
	currentState proto.State
	events       chan action.Action
}

func New() *kakiClient {
	return &kakiClient{
		log:          log.New(os.Stdout, "kaki: ", log.LstdFlags|log.Lmsgprefix),
		currentState: proto.Handshaking,
		events:       make(chan action.Action),
	}
}

func (k *kakiClient) stop() {
	close(k.events)
}

func (k *kakiClient) Process(ctx context.Context, packets <-chan []byte) chan action.Action {
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				k.log.Printf("client %s: Process stopping from context\n", ctx.Value("client_id"))
				break loop
			case pkt, ok := <-packets:
				if !ok {
					k.log.Printf("client %s: Process stopping on packets chan close\n", ctx.Value("client_id"))
					break loop
				}
				id, data, err := destructurePacket(pkt)
				if err != nil {
					k.dispatch(proto.EventFatalError{err})
				}
				f := handlerFuncFor(k.currentState, id)
				f(k, data)
			}
		}
		k.stop()
		k.log.Println("goodbye")
	}()
	return k.events
}

func (k *kakiClient) PacketFor(c event.Event) ([]byte, error) {
	switch cback := c.(type) {
	case event.Disconnect:
		return encodeDisconnect(k, cback)
	case event.AuthSuccess:
		return encodeAuthSuccess(k, cback)
	default:
		return nil, proto.ErrorUnsupportedCallback
	}
}

func (k *kakiClient) dispatch(e action.Action) {
	k.events <- e
}

func destructurePacket(pkt []byte) (id int32, data []byte, err error) {
	br := bytes.NewReader(pkt)
	id, err = decode.VarInt(br)
	if err != nil {
		return
	}
	dataidx := len(pkt) - br.Len()
	data = pkt[dataidx:]
	return
}
