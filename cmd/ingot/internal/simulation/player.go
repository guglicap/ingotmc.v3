package simulation

import (
	"fmt"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/mc"
	"github.com/guglicap/ingotmc.v3/world"
	"sync"
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
		sim.log.Println(err)
		return
	}
	player := sim.LoadPlayer(plInfo, cl)

	err = cl.ProcessEvent(event.NewPlayer{
		EID:            player.EID,
		Dim:            player.Dim,
		Gamemode:       player.Gamemode,
		RenderDistance: sim.renderDistance,
	})

	if err != nil {
		sim.log.Println(err)
		return
	}

	wg := sync.WaitGroup{}
	loadChunk := func(c world.ChunkCoords, cl Client) {
		wg.Add(1)
		defer wg.Done()
		sim.world.LoadChunk(player.Dim, c)
		ev := event.ChunkLoad{Dimension: player.Dim, Coords: c}
		ev.SetTriggeredBy(player.Name)
		err = cl.ProcessEvent(ev)
		if err != nil {
			sim.log.Println(err)
		}
	}
	chunkCoords := player.Pos.GetChunkCoords().WithinRadialDistance(int32(sim.renderDistance))
	for _, c := range chunkCoords {
		go loadChunk(c, cl)
	}
	wg.Wait()
	fmt.Println("send chunk done")
	err = cl.ProcessEvent(event.PlayerMoved{To: player.Pos})
	if err != nil {
		sim.log.Println(err)
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
