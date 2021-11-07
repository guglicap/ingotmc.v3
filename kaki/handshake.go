package kaki

import (
	"bytes"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/decode"
)

const (
	handshake_handshakingID int32 = 0x00
)

func handshakeFuncFor(id int32) handlerFunc {
	switch id {
	case handshake_handshakingID:
		return handleHandshake
	default:
		return func(k *kakiClient, _ []byte) {
			k.dispatch(proto.ErrorUnsupportedPacket(proto.Handshaking, id))
		}
	}
}

func handleHandshake(k *kakiClient, pkt []byte) {
	br := bytes.NewReader(pkt)
	version, err := decode.VarInt(br)
	if err != nil {
		k.dispatch(proto.EventFatalError{err })
		return
	}
	decode.String(br) // server addr, ignored
	decode.UShort(br) // server port, ignored
	nextState, err := decode.VarInt(br)
	if err != nil {
		k.dispatch(proto.EventFatalError{err })
		return
	}
	k.currentState = proto.State(nextState)
	if version != kakiVersion {
		k.dispatch(proto.ErrorMismatchedProtocol(version, kakiVersion))
		return
	}
}
