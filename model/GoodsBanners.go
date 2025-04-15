package model

import (
	"gorm.io/gorm"
	"time"
)

type GoodsBanners struct {
	Id        uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	GoodsId   int32          `gorm:"column:goods_id;type:int;comment:商品的id;not null;" json:"goods_id"`                  // 商品的id
	Title     string         `gorm:"column:title;type:varchar(225);comment:''轮播图标题'';not null;" json:"title"`           // ''轮播图标题''
	ImageUrl  string         `gorm:"column:image_url;type:varchar(225);comment:''图片URL地址'';not null;" json:"image_url"` // ''图片URL地址''
	LinkUrl   string         `gorm:"column:link_url;type:varchar(225);comment:''点击跳转链接'';not null;" json:"link_url"`    // ''点击跳转链接''
	Sort      int64          `gorm:"column:sort;type:bigint;comment:''排序权重值'';not null;" json:"sort"`                   // ''排序权重值''
	IsActive  int64          `gorm:"column:is_active;type:bigint;comment:''是否启用状态'';not null;" json:"is_active"`        // ''是否启用状态''
}

func (g *GoodsBanners) TableName() string {
	return "GoodsBanners"
}
