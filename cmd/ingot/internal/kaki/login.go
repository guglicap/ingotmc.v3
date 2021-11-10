package kaki

import (
	"bytes"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/decode"
)

func loginFuncFor(id int32) decodeFunc {
	switch id {
	case 0x00:
		return handleLoginStart
	default:
		return func(k *kakiClient, _ []byte) (action.Action, error) {
			return nil, proto.ErrorUnsupportedPacket(proto.Login, id)
		}
	}
}

func handleLoginStart(k *kakiClient, pkt []byte) (action.Action, error) {
	br := bytes.NewReader(pkt)
	name, err := decode.String(br)
	if err != nil {
		return nil, err
	}
	return action.NewConnection{Username: name}, nil
}
