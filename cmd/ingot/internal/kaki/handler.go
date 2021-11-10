package kaki

import (
	"github.com/guglicap/ingotmc.v3/proto"
)

type handlerFunc func(client *kakiClient, pkt []byte)

func handlerFuncFor(s proto.State, id int32) handlerFunc {
	switch s {
	case proto.Handshaking:
		return handshakeFuncFor(id)
	case proto.Login:
		return loginFuncFor(id)
	default:
		return func(k *kakiClient, _ []byte) {
			k.events <- proto.ErrorUnsupportedState(s)
		}
	}
}
