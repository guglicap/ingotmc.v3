package kaki

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/bearmini/bitstream-go"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/encode"
	"github.com/guglicap/ingotmc.v3/world"
	"github.com/guglicap/ingotmc.v3/world/block"
	"github.com/ingotmc/nbt"
	"io"
	"math"
)

const (
	play_ChunkData int32 = 0x22
)

type serializeError error

func serializeSafe(err error) {
	if err != nil {
		panic(serializeError(err))
	}
}

func recoverSerializeErr(p interface{}) error {
	if err, ok := p.(serializeError); ok {
		return err
	}
	panic(p)
}

func encodeChunkLoad(k *kakiClient, cL event.ChunkLoad) (pkt []byte, err error) {
	if !k.assertState(proto.Play) {
		err = proto.ErrorUnsupportedPacket(k.currentState, play_ChunkData)
		return
	}
	buf := &bytes.Buffer{}
	defer func() {
		x := recover()
		if x == nil {
			return
		}
		err = recoverSerializeErr(x)
	}()

	// encode position
	serializeSafe(encode.Int(cL.Coords.X, buf))
	serializeSafe(encode.Int(cL.Coords.Z, buf))

	serializeSafe(encode.Bool(true, buf)) // full chunk

	chData, err := k.world.ChunkDataFor(cL.Dimension, cL.Coords) // get chunk data for bitmask
	serializeSafe(err)
	serializeSafe(encode.VarInt(chData.GetBitMask(), buf))

	hMap, err := k.world.HeightMapFor(cL.Dimension, cL.Coords) // encode heightmap
	serializeSafe(err)
	hMapBytes, err := serializeHeightMap(hMap)
	serializeSafe(err)
	serializeSafe(nbt.Encode(nbt.Compound{
		"MOTION_BLOCKING": hMapBytes,
	}, buf))

	bI, err := k.world.BiomeDataFor(cL.Dimension, cL.Coords) // encode biomes
	serializeSafe(err)
	for _, b := range bI {
		serializeSafe(encode.Int(int32(b), buf))
	}

	data := &bytes.Buffer{}
	for i := range chData {
		if chData[i] == nil {
			continue
		}
		serializeSafe(serializeChunkSection(k.palette, *chData[i], data))
	}

	serializeSafe(encode.VarInt(int32(data.Len()), buf)) // data size
	_, err = io.Copy(buf, data)                          // data
	serializeSafe(err)
	serializeSafe(encode.VarInt(int32(0), buf)) // n block entities
	return
}

func serializeHeightMap(hMap world.HeightMap) (data []int64, err error) {
	buf := &bytes.Buffer{}
	bw := bitstream.NewWriter(buf)
	// note: copied from old ingot code, have no clue how it works rn
	for z := int32(0); z < 16; z++ {
		for x := int32(0); x < 16; x++ {
			bw.WriteNBitsOfUint16BE(9, hMap.HeightAt(x, z))
		}
	}
	if buf.Len() != 64/8*36 {
		err = errors.New("heighmap buf len isn't what expected")
		return
	}
	b := buf.Bytes()
	data = make([]int64, 36)
	for i := range data {
		data[i] = int64(binary.BigEndian.Uint64(b[i*8 : (i*8)+8]))
	}
	return
}

func serializeChunkSection(gp proto.GlobalPalette, cs proto.ChunkSection, w io.Writer) error {
	nonAir := 0
	unique := 0
	palette := make(map[block.Block]int)
	for _, b := range cs {
		if b != block.Air {
			nonAir++
		}
		if _, ok := palette[b]; ok {
			continue // if block already in palette
		}
		// this is an hack! we're setting an int for each block
		// when we initialize the actual palette array, this will be the index of this block
		// this works because unique starts at 0 and increases every time there's a new block state
		palette[b] = unique
		unique++
	}
	bpb := int(math.Ceil(math.Log2(float64(unique))))
	if bpb > 8 {
		return serializeChunkSectionDirect(nonAir, gp, cs, w)
	}

	pktPalette := make([]int32, unique)
	for bl, idx := range palette {
		// hack as above
		pktPalette[idx] = gp.IDFor(bl)
	}

	serializeSafe(encode.Short(int16(nonAir), w))
	serializeSafe(encode.UByte(uint8(bpb), w))
	serializeSafe(encode.VarInt(int32(len(pktPalette)), w))
	for _, v := range pktPalette {
		serializeSafe(encode.VarInt(v, w))
	}
	numOfLongs := bpb * 4096 / 8

	bitstreamBuf := make([]byte, 0, numOfLongs*8)
	bw := bitstream.NewWriter(bytes.NewBuffer(bitstreamBuf))
	for _, bl := range cs {
		idx := palette[bl]
		serializeSafe(bw.WriteNBitsOfUint32BE(uint8(bpb), uint32(idx)))
	}
	serializeSafe(bw.Flush())

	serializeSafe(encode.VarInt(int32(numOfLongs), w))
	for i := 0; i < numOfLongs; i++ {
		serializeSafe(
			encode.Long(
				int64(binary.BigEndian.Uint64(bitstreamBuf[i*8:(i+1)*8])),
				w,
			),
		)
	}
	return nil
}

func serializeChunkSectionDirect(nonAir int, palette proto.GlobalPalette, cs proto.ChunkSection, w io.Writer) error {
	serializeSafe(encode.Short(int16(nonAir), w))
	bpb := palette.BitsPerBlock()
	serializeSafe(encode.UByte(bpb, w))
	numOfLongs := int(bpb) * 4096 / 64
	bitstreamBuf := make([]byte, 0, numOfLongs*8)
	bw := bitstream.NewWriter(bytes.NewBuffer(bitstreamBuf))
	for _, bl := range cs {
		id := palette.IDFor(bl)
		serializeSafe(bw.WriteNBitsOfUint32BE(bpb, uint32(id)))
	}
	serializeSafe(bw.Flush())
	serializeSafe(encode.VarInt(int32(numOfLongs), w))
	for i := 0; i < numOfLongs; i++ {
		serializeSafe(
			encode.Long(
				int64(binary.BigEndian.Uint64(bitstreamBuf[i*8:(i+1)*8])),
				w,
			),
		)
	}
	return nil
}
