package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Redis struct {
	AdminDB *redis.Client
	UserDB  *redis.Client
}

func ConnectToRedis() *Redis {
	conn_str := viper.GetString("redis")
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
