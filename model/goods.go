package model

import (
	"gorm.io/gorm"
	"time"
)

// Goods 商品表
type Goods struct {
	GoodId       uint64         `gorm:"column:good_id;type:bigint UNSIGNED;primaryKey;not null;" json:"good_id"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	GoodName     string         `gorm:"column:good_name;type:varchar(30);comment:''商品的名称'';not null;" json:"good_name"`          // ''商品的名称''
	GoodPrice    float64        `gorm:"column:good_price;type:decimal(10, 2);comment:''商品的价格'';not null;" json:"good_price"`     // ''商品的价格''
	ParentId     int8           `gorm:"column:parent_id;type:tinyint;comment:''分类ID'';not null;" json:"parent_id"`               // ''分类ID''
	Image        string         `gorm:"column:image;type:varchar(225);comment:''商品的主图片'';not null;" json:"image"`                // ''商品的主图片''
	Feature      string         `gorm:"column:feature;type:varchar(225);comment:''商品的卖点'';not null;" json:"feature"`             // ''商品的卖点''
	GoodImage    string         `gorm:"column:good_image;type:varchar(225);comment:''商品主图'';not null;" json:"good_image"`        // ''商品主图''
	GoodsImages  string         `gorm:"column:goods_images;type:varchar(225);comment:''商品轮播图主图'';not null;" json:"goods_images"` // ''商品轮播图主图''
	GoodVideo    string         `gorm:"column:good_video;type:varchar(225);comment:''商品的视频'';not null;" json:"good_video"`       // ''商品的视频''
	GoodsPoster  string         `gorm:"column:goods_poster;type:varchar(225);comment:''商品海报'';not null;" json:"goods_poster"`    // ''商品海报''
	GoodBrand    string         `gorm:"column:good_brand;type:varchar(100);comment:''商品品牌'';not null;" json:"good_brand"`        // ''商品品牌''
	GoodSupplier string         `gorm:"column:good_supplier;type:varchar(100);comment:''供货商'';not null;" json:"good_supplier"`   // ''供货商''
	GoodSales    int64          `gorm:"column:good_sales;type:bigint;comment:''销售量'';not null;" json:"good_sales"`               // ''销售量''
	GoodViews    int64          `gorm:"column:good_views;type:bigint;comment:''商品的浏览量'';not null;" json:"good_views"`            // ''商品的浏览量''
}

func (g *Goods) TableName() string {
	return "goods"
}

type GoodsCategories struct {
	Id           uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(20);comment:分类名称;not null;" json:"name"`                      // 分类名称
	Image        string    `gorm:"column:image;type:varchar(255);comment:分类图片URL;default:NULL;" json:"image"`            // 分类图片URL
	ParentId     uint32    `gorm:"column:parent_id;type:int UNSIGNED;comment:父级分类ID;default:NULL;" json:"parent_id"`     // 父级分类ID
	Sort         int32     `gorm:"column:sort;type:int;comment:排序值;not null;default:0;" json:"sort"`                     // 排序值
	IsVisible    int8      `gorm:"column:is_visible;type:tinyint(1);comment:是否显示（1=是，0=否）;default:1;" json:"is_visible"` // 是否显示（1=是，0=否）
	Status       int8      `gorm:"column:status;type:tinyint(1);comment:状态（1=正常，2=禁用）;default:1;" json:"status"`         // 状态（1=正常，2=禁用）
	ProductCount int32     `gorm:"column:product_count;type:int;comment:商品数量;default:0;" json:"product_count"`           // 商品数量
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
}

func (g *GoodsCategories) TableName() string {
	return "goods_categories"
}

type GoodSeckills struct {
	SeckillsId uint64         `gorm:"column:seckills_id;type:bigint UNSIGNED;primaryKey;not null;" json:"seckills_id"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	GoodName   string         `gorm:"column:good_name;type:varchar(30);comment:''商品的名称'';not null;" json:"good_name"` // ''商品的名称''
	GoodId     int32          `gorm:"column:good_id;type:int;comment:''参与秒杀的商品'';not null;" json:"good_id"`           // ''参与秒杀的商品''
}

func (g *GoodSeckills) TableName() string {
	return "good_seckills"
}

type GoodSeckillTimes struct {
	SeckillTimeId uint64    `gorm:"column:seckill_time_id;type:bigint UNSIGNED;primaryKey;not null;" json:"seckill_time_id"`
	StartTime     time.Time `gorm:"column:start_time;type:datetime(3);comment:''秒杀开始时间'';default:NULL;" json:"start_time"` // ''秒杀开始时间''
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	EndTime       time.Time `gorm:"column:end_time;type:datetime(3);comment:''秒杀结束时间'';default:NULL;" json:"end_time"`        // ''秒杀结束时间''
	SeckillId     int64     `gorm:"column:seckill_id;type:bigint;comment:''参与秒杀商品的ID'';not null;" json:"seckill_id"`          // ''参与秒杀商品的ID''
	SeckillPrice  float64   `gorm:"column:seckill_price;type:decimal(10, 2);comment:''秒杀价格'';not null;" json:"seckill_price"` // ''秒杀价格''
	SeckillStatus int64     `gorm:"column:seckill_status;type:bigint;comment:''秒杀是否开启'';not null;" json:"seckill_status"`     // ''秒杀是否开启''
	SeckillNum    int32     `gorm:"column:seckill_num;type:int;comment:''秒杀商品的数量'';not null;" json:"seckill_num"`             // ''秒杀商品的数量''
}

func (g *GoodSeckillTimes) TableName() string {
	return "good_seckill_times"
}
