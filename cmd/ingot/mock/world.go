package mock

import (
	"github.com/guglicap/ingotmc.v3/proto"
	"github.com/guglicap/ingotmc.v3/world"
	"github.com/guglicap/ingotmc.v3/world/block"
	"log"
)

type World struct {
}

func (w World) BlockAt(bCoords world.BlockCoords) block.Block {
	panic("implement me")
}

func (w World) SetBlockAt(cCoords world.ChunkCoords, block block.Block) {
	panic("implement me")
}

func (w World) ChunkAt(coords world.ChunkCoords) world.Chunk {
	panic("implement me")
}

func (w World) LoadChunk(dim world.Dimension, coords world.ChunkCoords) {
	log.Printf("worldmock: loading chunk %+v\n", coords)
}

func (w World) ChunkDataFor(dim world.Dimension, coords world.ChunkCoords) (proto.ChunkData, error) {
	chData := proto.ChunkData{}
	chData[0] = &proto.ChunkSection{}
	for i := range chData[0] {
		chData[0][i] = block.OakPlanks
	}
	return chData, nil
}

func (w World) HeightMapFor(dim world.Dimension, coords world.ChunkCoords) (world.HeightMap, error) {
	h := world.HeightMap{}
	for i := range h {
		h[i] = 16
	}
	return h, nil
}

func (w World) BiomeDataFor(dim world.Dimension, coords world.ChunkCoords) (proto.BiomeInfo, error) {
	return proto.AllVoidBiomeInfo(), nil
}

func (w World) Level() world.Level {
	return world.LevelFlat
}

func (w World) Seed() world.Seed {
	return "test"
}
