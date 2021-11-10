package kaki

import (
	"crypto/md5"
	"github.com/guglicap/ingotmc.v3/mc"
)

type offlineAuthenticator func(username string) (mc.UUID, error)

func (oA offlineAuthenticator) Authenticate(user string) (mc.UUID, error) {
	return oA(user)
}

var OfflineAuthenticator offlineAuthenticator = func(username string) (mc.UUID, error) {
	hash := md5.Sum([]byte("OfflineUser:" + username))
	hash[6] = hash[6]&0x0f | 0x30
	hash[8] = hash[8]&0x3f | 0x80
	return hash, nil
}
