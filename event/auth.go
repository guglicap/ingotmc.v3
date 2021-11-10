package event

import (
	"github.com/guglicap/ingotmc.v3/mc"
)

type AuthSuccess struct {
	connEvent
	UUID     mc.UUID
	Username string
}
