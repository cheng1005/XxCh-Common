package model

import (
	"gorm.io/gorm"
	"time"
)

// Goodsseckilltime 秒杀表
type Goodsseckilltime struct {
	Id               int32          `gorm:"column:id;type:int;comment:秒杀表的ID;primaryKey;not null;" json:"id"`                              // 秒杀表的ID
	GoodSeckillingId int32          `gorm:"column:good_seckilling_id;type:int;comment:参与秒杀的商品的Id;not null;" json:"good_seckilling_id"` // 参与秒杀的商品的Id
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建的时间;not null;" json:"created_at"`                 // 创建的时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(3);comment:修改的时间;not null;" json:"updated_at"`                 // 修改的时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除的时间;default:NULL;" json:"deleted_at"`             // 删除的时间
	StartTime        time.Time      `gorm:"column:start_time;type:datetime(3);comment:秒杀开始的时间;not null;" json:"start_time"`             // 秒杀开始的时间
	EndTime          time.Time      `gorm:"column:end_time;type:datetime(3);comment:秒杀结束的时间;not null;" json:"end_time"`                 // 秒杀结束的时间
}

func (g *Goodsseckilltime) TableName() string {
	return "GoodsSeckillTime"
}
