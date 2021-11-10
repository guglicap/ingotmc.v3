package proto

import "github.com/guglicap/ingotmc.v3/world"

// BiomeInfo holds information about all the biomes in the chunk
// This is a Y-Z-X indexed array.
type BiomeInfo [1024]world.Biome

// BiomeAt returns the biome for the given coords
func (bI BiomeInfo) BiomeAt(coords world.BlockCoords) world.Biome {
	relX, relY, relZ := coords.X&15, coords.Y&15, coords.Z&15
	x, y, z := relX&3, relY&3, relZ&3
	return bI[16*y+4*z+x]
}

func AllVoidBiomeInfo() (bI BiomeInfo) {
	for i := range bI {
		bI[i] = world.BiomeVoid
	}
	return
}
