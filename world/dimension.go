package world

// Dimension represents a world dimension.
// One of: nether (-1), overworld (0), end (1)
type Dimension int32

// String implements the Stringer interface.
// Returns "nether", "overworld" or "end" accordingly.
// Returns "<unknown??>" if no known dimension is passed.
func (dt Dimension) String() (s string) {
	switch dt {
	case Nether:
		s = "nether"
	case Overworld:
		s = "overworld"
	case TheEnd:
		s = "end"
	default:
		s = "<unknown??>"
	}
	return
}

const (
	Nether    Dimension = -1
	Overworld Dimension = 0
	TheEnd    Dimension = 1
)
