package inits

import (
	"context"
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/go-redis/redis/v8"
	"log"
)

func InitRedis() {
	redisConfig := global.Nacos.Redis
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})
	err := global.RDB.Ping(context.Background()).Err()
	if err != nil {
		return
	}
	log.Println("redis init success")
}
