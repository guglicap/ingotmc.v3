package kaki

import (
	"bytes"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/encode"
)

const (
	play_playerPosLook = 0x36
)

func encodePlayerMoved(k *kakiClient, ev event.PlayerMoved) (pkt []byte, err error) {
	if !k.assertState(proto.Play) {
		err = proto.ErrorUnsupportedPacket(k.currentState, play_playerPosLook)
		return
	}
	buf := bytes.NewBuffer(pkt)
	defer func() {
		err = recoverSerializeErr(recover())
	}()
	serializeSafe(encode.VarInt(play_playerPosLook, buf)) // id
	serializeSafe(encode.Double(ev.To.X, buf))
	serializeSafe(encode.Double(ev.To.Y, buf))
	serializeSafe(encode.Double(ev.To.Z, buf))
	serializeSafe(encode.Float(ev.Yaw, buf))
	serializeSafe(encode.Float(ev.Pitch, buf))
	serializeSafe(encode.UByte(0x00, buf))
	serializeSafe(encode.VarInt(0x1337, buf))
	pkt = buf.Bytes()
	return
}
