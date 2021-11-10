package kaki

import (
	"bytes"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/decode"
)

func loginFuncFor(id int32) handlerFunc {
	switch id {
	case 0x00:
		return handleLoginStart
	default:
		return func(k *kakiClient, _ []byte) {
			k.events <- proto.ErrorUnsupportedPacket(proto.Login, id)
		}
	}
}

func handleLoginStart(k *kakiClient, pkt []byte) {
	br := bytes.NewReader(pkt)
	name, err := decode.String(br)
	if err != nil {
		k.dispatch(proto.EventFatalError{err})
		return
	}
	k.dispatch(action.NewConnection{name})
}
