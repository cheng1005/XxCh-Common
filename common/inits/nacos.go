package inits

import (
	"encoding/json"
	"fmt"
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go.uber.org/zap"
)

func InitNaCos() {
	nacosConfig := global.Config.NacosConfig
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(nacosConfig.Port),
			Scheme:      "http",
		},
	}
	// Another way of create config client for dynamic configuration (recommend)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return
	}
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(config), &global.Nacos)
	if err != nil {
		zap.L().Error(fmt.Sprintf("json Unmarshal faild,%v", err))
	}

}
