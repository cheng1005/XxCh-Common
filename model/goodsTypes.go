package model

import (
	"gorm.io/gorm"
	"time"
)

// Goodstypes 商品分类表
type Goodstypes struct {
	Id                  uint32         `gorm:"column:id;type:int UNSIGNED;comment:分类的ID;primaryKey;not null;" json:"id"`                                              // 分类的ID
	CategoryName        string         `gorm:"column:category_name;type:varchar(10);comment:分类名称1.女装类,2.男装类，3.食品类，4.数码类;not null;" json:"category_name"` // 分类名称1.女装类,2.男装类，3.食品类，4.数码类
	CategoryUrl         string         `gorm:"column:category_url;type:varchar(255);comment:分类的图片地址;not null;" json:"category_url"`                               // 分类的图片地址
	CategoryPId         uint32         `gorm:"column:category_p_id;type:int UNSIGNED;comment:分类的父级id;not null;default:0;" json:"category_p_id"`                     // 分类的父级id
	CategoryDescription string         `gorm:"column:category_description;type:text;comment:商品分类描述;" json:"category_description"`                                  // 商品分类描述
	CreatedAt           time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`                                          // 创建时间
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                                          // 修改时间
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;default:NULL;" json:"deleted_at"`                                         // 删除时间
	CategoryRecommend   int8           `gorm:"column:category_recommend;type:tinyint(1);comment:是否推荐;not null;" json:"category_recommend"`                           // 是否推荐
	SortOrder           uint32         `gorm:"column:sort_order;type:int UNSIGNED;comment:分类排序;default:0;" json:"sort_order"`                                        // 分类排序
	CategoryDisplay     int8           `gorm:"column:category_display;type:tinyint(1);comment:是否显示;not null;" json:"category_display"`                               // 是否显示
	Children            []Goodstypes   `gorm:"foreignKey:category_p_id" json:"goodstypes"`
}

func (g *Goodstypes) TableName() string {
	return "GoodsTypes"
}
