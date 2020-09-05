package log

import (
	"testing"

	"gin-frame/pkg/cfg"
)

var (
	// _ = g.Config().
	_ = cfg.Init("../../config/config.yml")
)

func TestDebug(t *testing.T) {
	Debug("123123")
}
