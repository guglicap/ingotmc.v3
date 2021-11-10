package proto

import "github.com/guglicap/ingotmc.v3/world"

// ChunkData groups all the 16 vertical ChunkSections this chunk is composed of.
type ChunkData [16]*ChunkSection

// GetBitMask generates an int32 with bits set to 1 for every non-nil section of ChunkData.
// LSB represents the section from y = 0 to y = 15.
func (c ChunkData) GetBitMask() int32 {
	bitMask := uint32(0x00)
	for i, cs := range c {
		if cs == nil {
			continue
		}
		apply := uint32(0x01) << i
		bitMask |= apply
	}
	return int32(bitMask)
}

// ChunkProvider provides access to chunk information for protocol purposes
type ChunkProvider interface {
	// ChunkDataFor returns ChunkData the chunk at the given ChunkCoords in Dimension dim
	ChunkDataFor(dim world.Dimension, coords world.ChunkCoords) (ChunkData, error)

	// HeightMapFor returns the HeightMap for the chunk at the given ChunkCoords in Dimension dim
	HeightMapFor(dim world.Dimension, coords world.ChunkCoords) (world.HeightMap, error)

	// BiomeDataFor returns BiomeData for the chunk at the given ChunkCoords in Dimension dim
	BiomeDataFor(dim world.Dimension, coords world.ChunkCoords) (BiomeInfo, error)
}
