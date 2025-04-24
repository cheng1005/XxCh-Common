package model

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	Id               uint64         `gorm:"column:id;type:bigint UNSIGNED;comment:编号;primaryKey;not null;" json:"id"`                               // 编号
	GoodName         string         `gorm:"column:good_name;type:varchar(50);comment:商品名称;not null;" json:"good_name"`                              // 商品名称
	GoodDescribe     string         `gorm:"column:good_describe;type:varchar(50);comment:商品描述;not null;" json:"good_describe"`                      // 商品描述
	GoodPrice        uint32         `gorm:"column:good_price;type:decimal(10, 2) UNSIGNED;comment:商品价格;not null;" json:"good_price"`                // 商品价格
	GoodNum          uint64         `gorm:"column:good_num;type:bigint UNSIGNED;comment:商品数量;not null;" json:"good_num"`                            // 商品数量
	ReceivingAddress string         `gorm:"column:receiving_address;type:varchar(50);comment:收货地址;not null;" json:"receiving_address"`              // 收货地址
	BuyerMessage     string         `gorm:"column:buyer_message;type:varchar(50);comment:买家留言;not null;" json:"buyer_message"`                      // 买家留言
	DeliveryMethod   uint64         `gorm:"column:delivery_method;type:bigint UNSIGNED;comment:配送方式 1快递 2自提;not null;" json:"delivery_method"`      // 配送方式 1快递 2自提
	FreightAmount    uint32         `gorm:"column:freight_amount;type:decimal(10, 2) UNSIGNED;comment:运费金额;not null;" json:"freight_amount"`        // 运费金额
	Coupon           string         `gorm:"column:coupon;type:varchar(50);comment:优惠券;not null;" json:"coupon"`                                     // 优惠券
	Name             string         `gorm:"column:name;type:varchar(20);comment:收货人;not null;" json:"name"`                                         // 收货人
	Points           string         `gorm:"column:points;type:varchar(255);comment:积分抵扣;not null;" json:"points"`                                   // 积分抵扣
	Phone            string         `gorm:"column:phone;type:char(11);comment:手机号;not null;" json:"phone"`                                          // 手机号
	PaymentStatus    uint64         `gorm:"column:payment_status;type:bigint UNSIGNED;comment:支付状态 1未支付 2待支付 3已支付;not null;" json:"payment_status"` // 支付状态 1未支付 2待支付 3已支付
	OrderSource      string         `gorm:"column:order_source;type:varchar(255);comment:订单来源;not null;" json:"order_source"`                       // 订单来源
	OrderType        string         `gorm:"column:order_type;type:varchar(255);comment:订单类型;not null;" json:"order_type"`                           // 订单类型
	PayPrice         uint32         `gorm:"column:pay_price;type:decimal(10, 2) UNSIGNED;comment:实付金额;not null;" json:"pay_price"`                  // 实付金额
	OrderId          uint64         `gorm:"column:order_id;type:bigint UNSIGNED;comment:订单编号;not null;" json:"order_id"`                            // 订单编号
	TimeOfPayment    string         `gorm:"column:time_of_payment;type:varchar(50);comment:支付时间;not null;" json:"time_of_payment"`                  // 支付时间
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (o *Orders) TableName() string {
	return "orders"
}
