package event

import "github.com/guglicap/ingotmc.v3/world"

type PlayerMoved struct {
	defaultEvent
	To         world.Coords
	Yaw, Pitch float32
}
