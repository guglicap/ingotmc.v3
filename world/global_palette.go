package world

import "github.com/guglicap/ingotmc.v3/world/block"

// GlobalPalette maps Blocks to their minecraft ID
type GlobalPalette interface {
	// IDFor returns the numerical ID for the given block
	IDFor(block.Block) int32

	// NamespacedIDFor returns the namespaced ID (i.e. "minecraft:air") for the given block
	NamespacedIDFor(block.Block) string
}
