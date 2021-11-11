package mock

import "github.com/guglicap/ingotmc.v3/world/block"

type Palette struct{}

func (p Palette) IDFor(bl block.Block) (id int32) {
	switch bl {
	case block.Air:
		id = 0
	case block.OakPlanks:
		id = block.OakPlanks
	default:
		id = 1
	}
	return
}

func (p Palette) NamespacedIDFor(block block.Block) string {
	panic("implement me")
}

func (p Palette) BitsPerBlock() uint8 {
	return 14
}
