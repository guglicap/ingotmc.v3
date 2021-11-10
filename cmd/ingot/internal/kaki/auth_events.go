package kaki

import (
	"bytes"
	"fmt"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/encode"
)

func encodeAuthSuccess(k *kakiClient, as event.AuthSuccess) ([]byte, error) {
	buf := &bytes.Buffer{}
	if k.currentState != proto.Login {
		return nil, fmt.Errorf("cannot send authSuccess while in state %s", k.currentState)
	}
	k.currentState = proto.Play
	id := int32(0x02) // login success
	encode.VarInt(id, buf)
	encode.String(as.UUID.String(), buf)
	encode.String(as.Username, buf)
	return buf.Bytes(), nil
}
