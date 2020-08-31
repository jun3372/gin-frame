package user

import (
	"testing"

	"frame/pkg/cfg"
	"frame/pkg/log"
)

var (
	_      = cfg.Init("../../../config/config.yml")
	uCache = NewUserCache()
)

func TestCache_GetUserBaseCache(t *testing.T) {
	user, err := uCache.GetUserBaseCache(uint64(1))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("user=", user)
}

func TestCache_MultiGetUserBaseCache(t *testing.T) {
	uids := []uint64{1}

	result, err := uCache.MultiGetUserBaseCache(uids)
	if err != nil {
		t.Fatal(err)
	}

	log.Debug("result=", result[0])
}