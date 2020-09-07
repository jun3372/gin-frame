package redis

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"

	"github.com/jun3372/gin-frame/pkg/cfg"
	"github.com/jun3372/gin-frame/pkg/log"
)

var (
	_ = cfg.Init("../../config/config.yml")
)

func TestRedis(t *testing.T) {
	k := "order:poll"
	ctx := context.Background()
	// Redis().GeoAdd(ctx, k, &redis.GeoLocation{
	// 	Name:      "xxoo",
	// 	Longitude: 13.361389,
	// 	Latitude:  38.115556,
	// })

	radius := Redis().GeoRadius(ctx, k, 13.361389, 38.115556,
		&redis.GeoRadiusQuery{
			Radius:    1,
			Unit:      "mi",
			WithDist:  true,
			WithCoord: true,
		})

	j, _ := json.Marshal(radius.Val())
	log.Debug("radius=", cast.ToString(j))
}
