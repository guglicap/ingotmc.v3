package proto

import "github.com/guglicap/ingotmc.v3/world"

type GlobalPalette interface {
	world.GlobalPalette
	// BitsPerBlock returns the number of bits used to encode a block
	BitsPerBlock() uint8
}
