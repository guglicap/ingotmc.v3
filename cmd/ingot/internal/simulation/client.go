package simulation

import (
	"errors"
	"fmt"
	"github.com/guglicap/ingotmc.v3/action"
	"github.com/guglicap/ingotmc.v3/event"
	"github.com/guglicap/ingotmc.v3/mc"
)

type clientManager struct {
	clients map[mc.UUID]Client
}

type playerInfo struct {
	name string
	uuid mc.UUID
}

func (cm *clientManager) waitAuth(cl Client) (playerInfo, error) {
	actions := cl.Actions()
	ev, ok := <-actions
	fmt.Println("got action", ev)
	if !ok {
		return playerInfo{}, errors.New("problem with actions chan")
	}
	np, ok := ev.(action.NewPlayer)
	if !ok {
		err := errors.New("first event was not newPlayer")
		disconnect(cl, err)
		return playerInfo{}, err
	}
	cm.clients[np.UUID] = cl
	return playerInfo{np.Username, np.UUID}, nil
}

// Client describes the ability to interact with the simulation
type Client interface {
	// ProcessEvent
	ProcessEvent(event event.Event) error
	Actions() <-chan action.Action
}

func disconnect(client Client, err error) {
	client.ProcessEvent(event.Disconnect{
		Reason: err,
	})
}
