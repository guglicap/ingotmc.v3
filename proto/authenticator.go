package proto

import "github.com/guglicap/ingotmc.v3/mc"

// TODO: probably doesn't belong here.

type Authenticator interface {
	Authenticate(username string) (mc.UUID, error)
}
