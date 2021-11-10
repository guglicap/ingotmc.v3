package mc

import (
	"github.com/guglicap/ingotmc.v3/world"
	"testing"
)

func TestChunkCoords_DistanceTo(t *testing.T) {
	type fields struct {
		X int64
		Z int64
	}
	type args struct {
		oc world.ChunkCoords
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"distance", fields{0, 0}, args{world.ChunkCoords{-3, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := world.ChunkCoords{
				X: tt.fields.X,
				Z: tt.fields.Z,
			}
			if got := cc.RadialDistance(tt.args.oc); got != tt.want {
				t.Errorf("RadialDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
