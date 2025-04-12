package inits

import (
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/spf13/viper"
	"log"
)

func InitViper() {
	viper.SetConfigFile("./dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(err)
	}
	log.Println("viper read success")
}
