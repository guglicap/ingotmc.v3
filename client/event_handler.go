package client

import (
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/world"
)

type EventHandler struct {
	*Client
}

func (p EventHandler) SendChunk(coords world.ChunkCoords) {
	pkt, err := p.proto.PacketFor(ch)
	if err != nil {
		p.log.Println("couldn't encode chunk", coords, ": ", err)
		return
	}
	p.cbound <- pkt
}

func (p EventHandler) ProcessEvent(event event.Event) error {
	panic("implement me")
}

func (p EventHandler) Actions() <-chan action.Action {
	return p.actions
}

