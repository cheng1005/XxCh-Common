package model

import (
	"gorm.io/gorm"
	"time"
)

type Goodsbrands struct {
	Id               int32          `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	BrandName        string         `gorm:"column:brand_name;type:varchar(255);comment:品牌名称;not null;" json:"brand_name"`                                       // 品牌名称
	BrandFirstLetter string         `gorm:"column:brand_first_letter;type:char(3);comment:品牌首字母;not null;" json:"brand_first_letter"`                          // 品牌首字母
	BrandImage       string         `gorm:"column:brand_image;type:varchar(255);comment:品牌图片，存储图片路径等信息;default:NULL;" json:"brand_image"`              // 品牌图片，存储图片路径等信息
	BrandSort        uint32         `gorm:"column:brand_sort;type:int UNSIGNED;comment:品牌排序;default:NULL;" json:"brand_sort"`                                   // 品牌排序
	DisplayStatus    int8           `gorm:"column:display_status;type:tinyint(1);comment:显示状态，1表示显示，0表示不显示;not null;default:1;" json:"display_status"` // 显示状态，1表示显示，0表示不显示
	BrandDescription string         `gorm:"column:brand_description;type:text;comment:品牌描述;" json:"brand_description"`                                          // 品牌描述
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`                 // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime;comment:修改时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`                 // 修改时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;default:NULL;" json:"deleted_at"`                                       // 删除时间
}

func (g *Goodsbrands) TableName() string {
	return "goodsbrands"
}
