package simulation

import (
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/mc"
	"github.com/guglicap/ingotmc.v3/world"
)

type Player struct {
	Client `json:"-"`

	Name string
	UUID mc.UUID
	EID  mc.EID

	Pos      world.Coords
	Dim      world.Dimension
	Gamemode world.Gamemode
}

func (sim *Simulation) SpawnPlayerFor(cl Client) {
	plInfo, err := sim.waitAuth(cl)
	if err != nil {
		return
	}
	player := sim.LoadPlayer(plInfo, cl)
	chunkCoords := player.Pos.ToChunkCoords().WithinRadialDistance(3)
	for _, c := range chunkCoords {
		sim.world.LoadChunk(c)
		ev := event.ChunkLoad{Coords: c}
		ev.SetTriggeredBy(player.Name)
		err = cl.ProcessEvent(ev)
		if err != nil {
			// TODO: logging
		}
	}
}

func (sim *Simulation) LoadPlayer(info playerInfo, cl Client) *Player {
	pl, err := sim.PlayerStorage.LoadPlayer(info.uuid)
	if err == nil {
		pl.Client = cl
		return &pl
	}
	return &Player{
		Client: cl,

		Name: info.name,
		UUID: info.uuid,

		Pos:      sim.spawnPos,
		EID:      mc.EIDFrom(info.uuid[:]),
		Gamemode: world.Creative,
		Dim:      world.Overworld,
	}
}
