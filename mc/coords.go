package mc

import (
	"fmt"
	"math"
)

// Vector3 is more of a math thing, don't know why it's here.
type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) Dot(x Vector3) float64 {
	return v.X*x.X + v.Y*x.Y + v.Z*x.Z
}

func (v Vector3) Add(o Vector3) Vector3 {
	return Vector3{
		v.X + o.X,
		v.Y + o.Y,
		v.Z + o.Z,
	}
}

// Coords is an (x,y,z) float vector.
type Coords Vector3

// ToChunkCoords returns the ChunkCoords of the chunk enclosing this position in space.
func (c Coords) ToChunkCoords() ChunkCoords {
	return ChunkCoords{
		int32(c.X) >> 4,
		int32(c.Z) >> 4,
	}
}

func (c Coords) ToBlockCoords() BlockCoords {
	return BlockCoords{
		X: int32(math.Ceil(c.X)),
		Y: int32(math.Ceil(c.Y)),
		Z: int32(math.Ceil(c.Z)),
	}
}

// BlockCoords is an (x,y,z) int64 vector.
type BlockCoords struct {
	X, Y, Z int32
}

// GetChunkCoords returns the ChunkCoords for the chunk the block is placed at.
func (bc BlockCoords) GetChunkCoords() ChunkCoords {
	return ChunkCoords{
		X: bc.X >> 4,
		Z: bc.Z >> 4,
	}
}

// ChunkCoords is an (x,z) vector.
type ChunkCoords struct {
	X, Z int32
}

func (cCoords ChunkCoords) String() string {
	return fmt.Sprintf("(%d, %d)", cCoords.X, cCoords.Z)
}

// RadialDistance returns the distance to another chunk.
// It is used to determine which chunk should be loaded given a render radius (distance).
// e.g. the following chunks ( o ) should be within RadialDistance 1 of the center chunk ( x )
// 		o	o	o
// 		o	x	o
// 		o	o	o
func (cCoords ChunkCoords) RadialDistance(oc ChunkCoords) int {
	return int(math.Min(
		math.Abs(float64(oc.X-cCoords.X)),
		math.Abs(float64(oc.Z-cCoords.Z)),
	))
}

func (cCoords ChunkCoords) AllWithinRadius(r int32) (res []ChunkCoords) {
	for x := cCoords.X - r; x <= cCoords.X+r; x++ {
		for z := cCoords.Z - r;  z <= cCoords.Z+r; z++ {
			res = append(res, ChunkCoords{
				X: x,
				Z: z,
			})
		}
	}
	return
}
