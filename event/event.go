// Package event defines events for the simulation.
package event

import (
	"github.com/guglicap/ingotmc.v3/world"
)

// Event is a simulation response to an Action.
// Anything that happens in the simulation that players should know about should generate a Event.
// TODO: define the interface better, maybe add details about the event generation?
type Event interface {
	At() world.Coords
	TriggeredBy() string
}

type defaultEvent struct {
	at          world.Coords
	triggeredBy string
}

func (d *defaultEvent) SetTriggeredBy(user string) {
	d.triggeredBy = user
}

func (d defaultEvent) At() world.Coords {
	return d.at
}

func (d defaultEvent) TriggeredBy() string {
	return d.triggeredBy
}
