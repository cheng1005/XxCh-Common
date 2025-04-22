package model

import (
	"gorm.io/gorm"
	"time"
)

type Shopcart struct {
	Id         uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:购物车id;primaryKey;not null;" json:"id"`                               // 购物车id
	UserId     uint64         `gorm:"column:user_id;type:bigint UNSIGNED;comment:用户id;not null;default:0;" json:"user_id"`                        // 用户id
	GoodsId    uint64         `gorm:"column:goods_id;type:bigint UNSIGNED;comment:商品id;not null;default:0;" json:"goods_id"`                      // 商品id
	Num        int64          `gorm:"column:num;type:bigint;comment:商品数量;not null;" json:"num"`                                                 // 商品数量
	TotalPrice uint64         `gorm:"column:total_price;type:bigint UNSIGNED;comment:总价格;not null;default:0;" json:"total_price"`                // 总价格
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间;not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (s *Shopcart) TableName() string {
	return "shopCart"
}
