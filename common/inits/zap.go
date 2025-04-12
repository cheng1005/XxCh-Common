package inits

import "go.uber.org/zap"

func InitZap() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"D:/gocode/src/zg4/rk16/xxfz.log"}
	build, err := config.Build()
	if err != nil {
		return
	}
	zap.ReplaceGlobals(build)
}
