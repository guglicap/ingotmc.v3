package event

import (
	"github.com/guglicap/ingotmc.v3/mc"
	"github.com/guglicap/ingotmc.v3/world"
)

type AuthSuccess struct {
	connEvent
	UUID     mc.UUID
	Username string
}

type NewPlayer struct {
	defaultEvent
	EID            mc.EID
	Gamemode       world.Gamemode
	RenderDistance int
	Dim            world.Dimension
}
