package kaki

import (
	"bytes"
	"encoding/binary"
	"github.com/bearmini/bitstream-go"
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/proto/encode"
	"github.com/guglicap/ingotmc.v3/world/block"
	"io"
	"math"
)

type chunkSecSerializer struct {
	bpb    uint8
	nonAir int16

	//chSec *proto.ChunkSection // copy is expensive?
	globalPalette proto.GlobalPalette
	palette       map[block.Block]int32
	pktPalette    []int32
}

func (css *chunkSecSerializer) initPalette(chSec *proto.ChunkSection) {
	css.nonAir = 0
	unique := 0
	css.palette = make(map[block.Block]int32)
	for _, b := range chSec {
		if b != block.Air {
			css.nonAir++
		}
		if _, ok := css.palette[b]; ok {
			continue // if block already in palette
		}
		// this is an hack! we're setting an int for each block
		// when we initialize the actual palette array, this will be the index of this block
		// this works because unique starts at 0 and increases every time there's a new block state
		css.palette[b] = int32(unique)
		unique++
	}
	css.bpb = uint8(math.Ceil(math.Log2(float64(unique))))
	if css.bpb < 4 {
		css.bpb = 4
	}
	if css.bpb > 8 {
		css.bpb = css.globalPalette.BitsPerBlock()
	}
	if css.isDirect() {
		return
	}
	css.pktPalette = make([]int32, unique)
	for bl, idx := range css.palette {
		// hack as above
		css.pktPalette[idx] = css.globalPalette.IDFor(bl)
	}
}

func (css chunkSecSerializer) isDirect() bool {
	return css.bpb > 8
}

func (css chunkSecSerializer) serialize(cs *proto.ChunkSection, w io.Writer) {
	css.initPalette(cs)
	serializeSafe(encode.Short(css.nonAir, w))
	serializeSafe(encode.UByte(css.bpb, w))
	if !css.isDirect() {
		serializeSafe(encode.VarInt(int32(len(css.pktPalette)), w))
		for _, v := range css.pktPalette {
			serializeSafe(encode.VarInt(v, w))
		}
	}
	numOfLongs := int(css.bpb) * 4096 / 8
	bitstreamBuf := make([]byte, 0, numOfLongs*8)
	bw := bitstream.NewWriter(bytes.NewBuffer(bitstreamBuf))
	for _, bl := range cs {
		var v int32
		if css.isDirect() {
			v = css.globalPalette.IDFor(bl)
		} else {
			v = css.palette[bl]
		}
		serializeSafe(bw.WriteNBitsOfUint32BE(css.bpb, uint32(v)))
	}
	serializeSafe(bw.Flush())
	serializeSafe(encode.VarInt(int32(numOfLongs), w))
	for i := 0; i < numOfLongs; i++ {
		longBytes := bitstreamBuf[i*8 : (i+1)*8]
		long := binary.BigEndian.Uint64(longBytes)
		serializeSafe(encode.Long(int64(long), w))
	}
}
