package world

import "github.com/guglicap/ingotmc.v3/world/block"

type InfoProvider interface {
	Level() Level
	Seed() Seed
}

type Provider interface {
	InfoProvider

	BlockAt(bCoords BlockCoords) block.Block
	SetBlockAt(cCoords ChunkCoords, block block.Block)

	ChunkProvider
	ChunkLoader
}
