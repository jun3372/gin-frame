package gorm

import (
	"testing"

	"github.com/jun3372/gin-frame/pkg/cfg"
)

func TestInitDB(t *testing.T) {
	var (
		err error
	)

	if err = cfg.Init("../../config/config.yml"); err != nil {
		t.Fatal(err)
	}

	if err = InitDB(); err != nil {
		t.Fatal(err)
	}

	t.Log("config=", config)
	t.Logf("链接成功: db=%v", GetDB())
}
