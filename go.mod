module github.com/guglicap/ingotmc.v3

go 1.18

require github.com/google/uuid v1.3.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	golang.org/x/sys v0.2.0 // indirect
)

replace github.com/cskr/pubsub => ./cmd/ingot/internal/pubsub

require (
	github.com/bearmini/bitstream-go v0.0.0-20190121230027-bec1c9ea0d3c // indirect
	github.com/cskr/pubsub v0.0.0
	github.com/fatih/color v1.13.0
	github.com/ingotmc/nbt v0.0.1 // indirect
	github.com/pkg/errors v0.8.0 // indirect
)
