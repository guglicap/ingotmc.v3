package simulation

import (
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/world"
)

// EventHandler describes the ability to interact with the simulation
type EventHandler interface {
	SendChunk(ch world.ChunkCoords)
	ProcessEvent(event event.Event) error
	Actions() <-chan action.Action
}
