package global

import (
	"github.com/cheng1005/XxCh-Common/common/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	RDB    *redis.Client
	Config *config.AppConfig
	Nacos  *config.Nacos
)
