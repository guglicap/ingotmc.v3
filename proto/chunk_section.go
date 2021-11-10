package proto

import (
	"github.com/guglicap/ingotmc.v3/world/block"
)

// ChunkSection represents a 16x16x16 set of YZX indexed blocks
type ChunkSection [4096]block.Block
