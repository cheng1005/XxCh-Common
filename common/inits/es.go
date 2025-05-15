package inits

import (
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func InitEs() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://14.103.235.216:9200",
		},
		// ...
	}
	global.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("es init success")

}
