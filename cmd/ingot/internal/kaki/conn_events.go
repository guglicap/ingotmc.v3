package kaki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/encode"
)

func encodeDisconnect(k *kakiClient, dc event.Disconnect) ([]byte, error) {
	pkt := &bytes.Buffer{}
	var id int32
	switch k.currentState {
	case proto.Login:
		id = 0x00
	case proto.Play:
		id = 0x1B
	default:
		return nil, fmt.Errorf("cannot send disconnect while in state %s", k.currentState)
	}
	err := encode.VarInt(id, pkt)
	if err != nil {
		return nil, err
	}
	pktReason := struct {
		Text string `json:"text"`
	}{
		dc.Reason.Error(),
	}
	reason, err := json.Marshal(pktReason)
	if err != nil {
		return nil, err
	}
	encode.String(string(reason), pkt)
	return pkt.Bytes(), nil
}
