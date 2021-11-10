// Package world provides an interface to the minecraft world.
// It shouldn't care about implementation, but should still be useful to deal with worlds.
// A non-programmer minecraft player should be able to get a sense for what this package does.
package world

// Block needs to be implemented.
// TODO: implement Block.
type Block int

// ChunkSections represents all chunk sections this chunk is composed of.
// This is protocol specific but sane enough to be here, might change in the future.
type ChunkSections [16]*[4096]Block

// Chunk describes an area of 16x16 Block
type Chunk interface {
	// BlockAt returns the block at the given BlockCoords.
	// This BlockCoords should be relative to the chunk, if they aren't, BlockAt is going to make them so.
	BlockAt(bCoords BlockCoords) (Block, error)
	// AllBlocks returns an
	AllBlocks() (ChunkSections, error)
}

// ChunkProvider provides access to Chunks.
type ChunkProvider interface {
	LoadChunk(coords ChunkCoords) // TODO: maybe error
	ChunkAt(coords ChunkCoords) Chunk
}
