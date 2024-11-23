package redis

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	AdminDB *redis.Client
	UserDB  *redis.Client
}

func ConnectToRedis(conn_str string) *Redis {
	return &Redis{
		AdminDB: redis.NewClient(&redis.Options{
			Addr: conn_str,
			DB:   0,
		}),
		UserDB: redis.NewClient(&redis.Options{
			Addr: conn_str,
			DB:   1,
		}),
	}
}
