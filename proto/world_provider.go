package proto

import "github.com/guglicap/ingotmc.v3/world"

type WorldProvider interface {
	ChunkProvider
	world.InfoProvider
}
