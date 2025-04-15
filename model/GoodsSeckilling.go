package model

import (
	"gorm.io/gorm"
	"time"
)

// Goodsseckilling 秒杀商品id
type Goodsseckilling struct {
	Id            int32          `gorm:"column:id;type:int;comment:秒杀商品ID;primaryKey;not null;" json:"id"`                                 // 秒杀商品ID
	GoodsId int32 `gorm:"column:goods_id;type:int;comment:参加秒杀商品的ID;not null;" json:"goods_id"` // 参加秒杀商品的ID
	CreatedTa     time.Time      `gorm:"column:created_ta;type:datetime(3);comment:创建时间;not null;" json:"created_ta"`                      // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                      // 修改时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                  // 删除时间
	SeckillPrice  float32        `gorm:"column:seckill_price;type:decimal(10, 2);comment:秒杀的价格;not null;" json:"seckill_price"`           // 秒杀的价格
	StockLimit    int32          `gorm:"column:stock_limit;type:int;comment:秒杀的库存;not null;" json:"stock_limit"`                          // 秒杀的库存
	SeckillStatus string         `gorm:"column:seckill_status;type:enum('1 是', '2 否');comment:是否秒杀状态;not null;" json:"seckill_status"` // 是否秒杀状态
}

func (g *Goodsseckilling) TableName() string {
	return `GoodsSeckilling`
}
