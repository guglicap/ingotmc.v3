package world

// BiomeInfo provides information about biomes
type BiomeInfo interface {
	// BiomeAt returns the biome for the given coords
	BiomeAt(coords BlockCoords)
}

// Biome needs to be implemented
// TODO: implement biome, should specify type, climate conditions, mob spawning rules...
type Biome int32

const (
	BiomeVoid Biome = 127
)
