package dao_redis

import (
	"context"
	"fmt"
	"github.com/cheng1005/XxCh-Common/common/global"
)

// KeyExists 判断购物车商品key是否存在
func KeyExists(key string) (bool, error) {
	count, err := global.RDB.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CartGoods 购物车商品
func CartGoods(key string) string {
	result, _ := global.RDB.Get(context.Background(), key).Result()

	return result
}

// SyncGoodsStock 同步商品库存
func SyncGoodsStock(goodsId, goodsStock int64) {
	//goods_stock:goods_id_1
	key := fmt.Sprintf("goods_stock:goods_id:%d", goodsId)
	for i := 0; i < int(goodsStock); i++ {
		global.RDB.RPush(context.Background(), key, goodsId)
	}
}

// GetGoodsStock 查询商品库存
func GetGoodsStock(goodsId int) int64 {
	key := fmt.Sprintf("goods_stock:goods_id:%d", goodsId)
	return global.RDB.LLen(context.Background(), key).Val()
}
