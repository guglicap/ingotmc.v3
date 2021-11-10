package action

import (
	"github.com/guglicap/ingotmc.v3/mc"
)

type NewConnection struct {
	Username string
}

type NewPlayer struct {
	UUID     mc.UUID
	Username string
}
