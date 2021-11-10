package world

type Provider interface {
	Dimension() Dimension
	Level() Level
	Seed() Seed

	BlockAt(bCoords BlockCoords) Block
	SetBlockAt(cCoords ChunkCoords, block Block)

	ChunkProvider
}
