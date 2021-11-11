package kaki

import (
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/proto"
)

type decodeFunc func(client *kakiClient, pkt []byte) (action.Action, error)

func decodeFuncFor(s proto.State, id int32) decodeFunc {
	switch s {
	case proto.Handshaking:
		return handshakeFuncFor(id)
	case proto.Login:
		return loginFuncFor(id)
	default:
		return func(k *kakiClient, _ []byte) (action.Action, error) {
			return nil, proto.ErrorUnsupportedPacket(k.currentState, id)
		}
	}
}
