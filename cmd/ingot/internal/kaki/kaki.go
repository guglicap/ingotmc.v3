package kaki

import (
	"bytes"
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
	world        proto.WorldProvider
	chunkSecSerializer
}

func New() *kakiClient {
	return &kakiClient{
		log:          log.New(os.Stdout, "kaki: ", log.LstdFlags|log.Lmsgprefix),
		currentState: proto.Handshaking,
	}
}

func (k *kakiClient) ActionFor(pkt []byte) (action.Action, error) {
	id, data, err := destructurePacket(pkt)
	if err != nil {
		return nil, err
	}
	f := decodeFuncFor(k.currentState, id)
	return f(k, data)
}

func (k *kakiClient) PacketFor(c event.Event) ([]byte, error) {
	switch cback := c.(type) {
	case event.Disconnect:
		return encodeDisconnect(k, cback)
	case event.AuthSuccess:
		return encodeAuthSuccess(k, cback)
	case event.ChunkLoad:
		return encodeChunkLoad(k, cback)
	default:
		return nil, proto.ErrorUnsupportedCallback
	}
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

func (k *kakiClient) assertState(state proto.State, others ...proto.State) bool {
	match := false
	match = state == k.currentState
	if match {
		return true
	}
	for _, s := range others {
		if match {
			break
		}
		match = s == k.currentState
	}
	return match
}
