package world

import "math"

// vector3 is more of a math thing, don't know why it's here.
type vector3 struct {
	X, Y, Z float64
}

func (v vector3) Dot(x vector3) float64 {
	return v.X*x.X + v.Y*x.Y + v.Z*x.Z
}

func (v vector3) Add(o vector3) vector3 {
	return vector3{
		v.X + o.X,
		v.Y + o.Y,
		v.Z + o.Z,
	}
}

// Coords is an (x,y,z) float vector.
type Coords vector3

// GetChunkCoords returns the ChunkCoords of the chunk enclosing this position in space.
func (c Coords) GetChunkCoords() ChunkCoords {
	return ChunkCoords{
		int32(c.X) >> 4,
		int32(c.Z) >> 4,
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

// WithinRadialDistance returns all ChunkCoords within a radius r of this one.
func (cCoords ChunkCoords) WithinRadialDistance(r int32) (c []ChunkCoords) {
	c = make([]ChunkCoords, 0, 2*r+1)
	for z := cCoords.Z - r; z <= cCoords.Z+r; z++ {
		for x := cCoords.X - r; x <= cCoords.X+r; x++ {
			c = append(c, ChunkCoords{
				X: x,
				Z: z,
			})
		}
	}
	return
}
