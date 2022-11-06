package block

import "github.com/guglicap/ingotmc.v3/world"

type simpleSolidBlock struct {
	id int
}

func (s simpleSolidBlock) ID() int {
	return s.id
}

func (s simpleSolidBlock) Solid() bool {
	return true
}

func (s simpleSolidBlock) Properties() struct{} {
	return struct{}{}
}

var (
	Air       world.Block = simpleSolidBlock{0}
	OakPlanks world.Block = simpleSolidBlock{15}
	CHANGEME world.Block = simpleSolidBlock{14}
)
