package simulation

import (
	"encoding/json"
	"github.com/guglicap/ingotmc.v3/mc"
	"os"
	"path"
)

type PlayerStorage interface {
	LoadPlayer(mc.UUID) (Player, error)
	SavePlayer(Player) error
}

type playerStore struct {
	rootPath string
}

func (p playerStore) LoadPlayer(id mc.UUID) (Player, error) {
	filePath := path.Join(p.rootPath, id.String()+".json")
	file, err := os.Open(filePath)
	if err != nil {
		return Player{}, err
	}
	defer file.Close()
	pl := Player{}
	err = json.NewDecoder(file).Decode(&pl)
	return pl, err
}

func (p playerStore) SavePlayer(player Player) error {
	file, err := os.Create(path.Join(p.rootPath, player.UUID.String()+".json"))
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(player)
}
