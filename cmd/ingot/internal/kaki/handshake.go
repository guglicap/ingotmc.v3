package kaki

import (
	"bytes"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/decode"
)

const (
	handshake_handshakingID int32 = 0x00
)

func handshakeFuncFor(id int32) decodeFunc {
	switch id {
	case handshake_handshakingID:
		return handleHandshake
	default:
		return func(k *kakiClient, _ []byte) (action.Action, error) {
			return nil, proto.ErrorUnsupportedPacket(proto.Handshaking, id)
		}
	}
}

func handleHandshake(k *kakiClient, pkt []byte) (action.Action, error) {
	br := bytes.NewReader(pkt)
	version, err := decode.VarInt(br)
	if err != nil {
		return nil, err
	}
	decode.String(br) // server addr, ignored
	decode.UShort(br) // server port, ignored
	nextState, err := decode.VarInt(br)
	if err != nil {
		return nil, err
	}
	k.currentState = proto.State(nextState)
	if version != kakiVersion {
		return nil, proto.ErrorMismatchedProtocol(version, kakiVersion)
	}
	// TODO: action.Noop ?
	return nil, nil
}
