package user

import (
	"testing"

	"github.com/jun3372/gin-frame/pkg/cfg"
)

var (
	_    = cfg.Init("../../../config/config.yml")
	repo = NewUserRepo()
)

func TestUserRepo_GetUserByID(t *testing.T) {
	user, err := repo.GetUserByID(uint64(9))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("user=", user)
}
