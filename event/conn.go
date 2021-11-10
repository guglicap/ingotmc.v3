package event

import (
	"github.com/guglicap/ingotmc.v3/world"
)

type connEvent struct{}

func (c connEvent) At() world.Coords {
	return world.Coords{}
}

func (c connEvent) TriggeredBy() string {
	return ""
}

type Disconnect struct {
	connEvent
	Reason error
}
