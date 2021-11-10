// Package world provides an interface to the minecraft world.
// It shouldn't care about implementation, but should still be useful to deal with worlds.
// A non-programmer minecraft player should be able to get a sense for what this package does.
package world

import "github.com/guglicap/ingotmc.v3/world/block"

// HeightMap holds information about the topmost block for each (x,z) block in the chunk.
// Note: this might be needed for purposes other than protocol implementation, so it's here, might have to move it later.
// TODO: implement HeightMap
type HeightMap [16 * 16]uint16

// HeightAt returns the y coordinate of the highest solid block at position x, z
func (h HeightMap) HeightAt(x, z int32) uint16 {
	return h[16*z&15+x&15]
}

// Chunk describes an area of 16x16 Block
type Chunk interface {
	// Dimension returns the Dimension this chunk is located in
	Dimension() Dimension

	// HeightMap returns the chunk heightmap
	HeightMap() HeightMap

	BiomeInfo

	// BlockAt returns the block at the given BlockCoords.
	// This BlockCoords should be relative to the chunk, if they aren't, BlockAt is going to make them so.
	BlockAt(bCoords BlockCoords) (block.Block, error)
}

// ChunkProvider provides access to Chunks.
type ChunkProvider interface {
	ChunkAt(coords ChunkCoords) Chunk
}

// ChunkLoader provides an interface to load Chunks.
type ChunkLoader interface {
	// LoadChunk instructs the ChunkLoader to load the chunk at ChunkCoords in Dimension dim
	LoadChunk(dim Dimension, coords ChunkCoords)
}
