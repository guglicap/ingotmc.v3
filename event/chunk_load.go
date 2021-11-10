package event

import "github.com/guglicap/ingotmc.v3/world"

// ChunkLoad is emitted when a new Chunk is loaded.
// Should be processed by sending the chunk data to the client.
type ChunkLoad struct {
	defaultEvent
	Coords world.ChunkCoords
}
