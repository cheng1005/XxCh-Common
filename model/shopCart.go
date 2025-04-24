package model

import (
	"gorm.io/gorm"
	"time"
)

// Cart 购物车表
type Cart struct {
	Id        uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	UserId    uint64         `gorm:"column:user_id;type:bigint UNSIGNED;not null;" json:"user_id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间;not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (c *Cart) TableName() string {
	return "cart"
}

// CartItem 购物车商品表
type CartItem struct {
	Id         uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CartId     uint64         `gorm:"column:cart_id;type:bigint UNSIGNED;comment:购物车id;not null;default:0;" json:"cart_id"`                     // 购物车id
	GoodsId    uint64         `gorm:"column:goods_id;type:bigint UNSIGNED;comment:商品id;not null;default:0;" json:"goods_id"`                    // 商品id
	TotalPrice uint64         `gorm:"column:total_price;type:bigint UNSIGNED;comment:总价格;not null;default:0;" json:"total_price"`               // 总价格
	Num        uint64         `gorm:"column:num;type:bigint UNSIGNED;comment:数量;not null;default:0;" json:"num"`                                // 数量
	Selected   int8           `gorm:"column:selected;type:tinyint(1);comment:是否选中 ;not null;default:1;" json:"selected"`                        // 是否选中
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间;not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"` // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (c *CartItem) TableName() string {
	return "cart_item"
}
