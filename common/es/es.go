package es

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/cheng1005/XxCh-Common/model"
	"github.com/cheng1005/XxCh-Common/proto/goods"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
	r  map[string]interface{}
)

// SyncGoodsToEs 同步商品到es
func SyncGoodsToEs(goodsList []model.Goods) {
	for _, good := range goodsList {
		wg.Add(1)

		go func(goods model.Goods) {
			defer wg.Done()

			// Build the request body.
			data, err := json.Marshal(good)
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "goods",
				DocumentID: strconv.Itoa(int(good.Id)),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), global.Es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

		}(good)
	}
	wg.Wait()
}

// GoodsEsSearch 商品es搜索
func GoodsEsSearch(kewWord string) (goodsList []*goods.GoodsList) {
	var buf bytes.Buffer
	query := map[string]interface{}{}
	if kewWord != "" {
		log.Println(kewWord)

		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"good_name": kewWord,
				},
			},
		}
	} else {
		log.Println(1)
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		}
		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query: %s", err)
		}
	}
	res, err := global.Es.Search(
		global.Es.Search.WithContext(context.Background()),
		global.Es.Search.WithIndex("goods"),
		global.Es.Search.WithBody(&buf),
		global.Es.Search.WithTrackTotalHits(true),
		global.Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		xx := hit.(map[string]interface{})["_source"].(map[string]interface{})
		goodsList = append(goodsList, &goods.GoodsList{
			GoodName:     xx["good_name"].(string),
			GoodPrice:    uint64(xx["good_price"].(float64)),
			ParentId:     uint64(xx["parent_id"].(float64)),
			GoodNum:      uint64(xx["good_num"].(float64)),
			GoodStatus:   uint64(xx["good_status"].(float64)),
			Image:        xx["image"].(string),
			Feature:      xx["feature"].(string),
			GoodImage:    xx["good_image"].(string),
			GoodVideo:    xx["good_video"].(string),
			GoodsPoster:  xx["goods_poster"].(string),
			GoodBrand:    xx["good_brand"].(string),
			GoodSupplier: xx["good_supplier"].(string),
			GoodSales:    uint64(xx["good_sales"].(float64)),
			GoodViews:    uint64(xx["good_views"].(float64)),
			ProductCode:  xx["product_code"].(string),
		})
	}
	return
}
