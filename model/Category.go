package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id        int32          `gorm:"column:id;type:int;comment:分类ID，主键;primaryKey;not null;" json:"id"`                                        // 分类ID，主键
	Name      string         `gorm:"column:Name;type:varchar(255);comment:分类名称;not null;" json:"Name"`                                         // 分类名称
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"` // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                          // 删除时间
}

func (c *Category) TableName() string {
	return "Category"
}
