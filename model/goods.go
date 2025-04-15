package model

import (
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	Id             uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:商品ID;primaryKey;not null;" json:"id"`                               // 商品ID
	ImageName      string         `gorm:"column:image_name;type:varchar(255);comment:图片地址;not null;" json:"image_name"`                             // 图片地址
	GoodsName      string         `gorm:"column:goods_name;type:varchar(255);comment:商品名称;not null;" json:"goods_name"`                             // 商品名称
	CategoryId     int64          `gorm:"column:category_id;type:bigint;comment:商品分类ID;not null;" json:"category_id"`                               // 商品分类ID
	Price          float32        `gorm:"column:price;type:decimal(10, 2) UNSIGNED ZEROFILL;comment:销售价格;not null;" json:"price"`                   // 销售价格
	Stock          uint32         `gorm:"column:stock;type:int(11) UNSIGNED ZEROFILL;comment:库存数量;not null;" json:"stock"`                          // 库存数量
	Status         uint8          `gorm:"column:status;type:tinyint(1) UNSIGNED ZEROFILL;comment:商品状态：0下架/1上架;not null;" json:"status"`             // 商品状态：0下架/1上架
	Description    string         `gorm:"column:description;type:text;comment:商品描述;" json:"description"`                                            // 商品描述
	BrandId        int64          `gorm:"column:brand_id;type:bigint;comment:品牌ID;not null;" json:"brand_id"`                                       // 品牌ID
	IsPromotion    uint8          `gorm:"column:is_promotion;type:tinyint(3) UNSIGNED ZEROFILL;comment:是否促销：0否/1是;not null;" json:"is_promotion"`   // 是否促销：0否/1是
	ProductionDate string         `gorm:"column:production_date;type:varchar(255);comment:生产日期;not null;" json:"production_date"`                   // 生产日期
	PromotionPrice float32        `gorm:"column:promotion_price;type:decimal(10, 2);comment:促销价格;not null;" json:"promotion_price"`                 // 促销价格
	IsSub          int8           `gorm:"column:is_sub;type:tinyint(1);comment:是否单独分佣;not null;default:0;" json:"is_sub"`                           // 是否单独分佣
	ExpiryDate     int32          `gorm:"column:expiry_date;type:int;comment:保质期天数;not null;" json:"expiry_date"`                                   // 保质期天数
	SalesVolume    int32          `gorm:"column:sales_volume;type:int;comment:商品销量;not null;" json:"sales_volume"`                                  // 商品销量
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"` // 修改时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (g *Goods) TableName() string {
	return "Goods"
}
