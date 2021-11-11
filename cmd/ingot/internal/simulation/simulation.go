package simulation

import (
	"github.com/guglicap/ingotmc.v3/mc"
	"github.com/guglicap/ingotmc.v3/world"
	"log"
	"os"
)

type Simulation struct {
	clientManager
	world world.Provider

	PlayerStorage

	log            *log.Logger
	spawnPos       world.Coords
	renderDistance int
}

func New(w world.Provider) *Simulation {
	return &Simulation{
		log:            log.New(os.Stdout, "sim: ", log.LstdFlags|log.Lmsgprefix),
		clientManager:  clientManager{clients: make(map[mc.UUID]Client)},
		PlayerStorage:  playerStore{rootPath: "srv/player"},
		spawnPos:       world.Coords{0.0, 16.0, 0.0},
		renderDistance: 16,
		world:          w,
	}
}
