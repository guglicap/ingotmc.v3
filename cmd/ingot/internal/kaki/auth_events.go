package kaki

import (
	"bytes"
	"fmt"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/encode"
)

const (
	play_JoinGame = 0x26
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

func encodeJoinGame(k *kakiClient, np event.NewPlayer) (pkt []byte, err error) {
	if !k.assertState(proto.Play) {
		err = proto.ErrorUnsupportedPacket(k.currentState, play_JoinGame)
		return
	}
	defer func() {
		err = recoverSerializeErr(recover())
	}()
	buf := bytes.NewBuffer(pkt)
	serializeSafe(encode.VarInt(play_JoinGame, buf))
	serializeSafe(encode.Int(int32(np.EID), buf))
	serializeSafe(encode.UByte(uint8(np.Gamemode), buf))
	serializeSafe(encode.Int(int32(np.Dim), buf))
	serializeSafe(encode.Long(int64(k.world.Seed().GetHash()), buf))
	encode.UByte(5, buf)
	serializeSafe(encode.String(string(k.world.Level()), buf))
	serializeSafe(encode.VarInt(int32(np.RenderDistance), buf))
	encode.Bool(false, buf)
	encode.Bool(false, buf)
	pkt = buf.Bytes()
	return
}
