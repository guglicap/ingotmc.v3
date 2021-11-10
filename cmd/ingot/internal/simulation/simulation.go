package simulation

import (
	"github.com/guglicap/ingotmc.v3/mc"
	"github.com/guglicap/ingotmc.v3/world"
)

type Simulation struct {
	clientManager
	world world.Provider

	PlayerStorage

	spawnPos       world.Coords
	renderDistance int
}

func New() *Simulation {
	return &Simulation{
		clientManager: clientManager{clients: make(map[mc.UUID]Client)},
		PlayerStorage: playerStore{rootPath: "srv/player"},
		spawnPos:      world.Coords{0.0, 0.0, 0.0},
	}
}
