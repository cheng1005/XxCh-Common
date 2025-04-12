package inits

import (
	"fmt"
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/hashicorp/consul/api"
	"log"
)

type Consul struct {
	consulClient *api.Client
}
type Register struct {
	ID      string
	Name    string
	Tags    []string
	Port    int
	Address string
}

func NewConsulClient(host string, port uint) (*Consul, error) {

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", host, port)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, nil
	}
	return &Consul{client}, err
}

func (c *Consul) RegisterConsul(req Register) error {

	registration := &api.AgentServiceRegistration{
		ID:      req.ID,
		Name:    req.Name,
		Tags:    req.Tags,
		Port:    req.Port,
		Address: req.Address,
	}
	return c.consulClient.Agent().ServiceRegister(registration)

}

func InitConsul() {
	consulConfig := global.Nacos.Consul
	client, err := NewConsulClient(consulConfig.Host, uint(consulConfig.Port))
	if err != nil {
		return
	}
	req := Register{
		ID:      "fz",
		Name:    consulConfig.Name,
		Tags:    nil,
		Port:    1112,
		Address: "localhost",
	}
	err = client.RegisterConsul(req)
	if err != nil {
		panic(err)
	}
	log.Println("consul register success")
}
