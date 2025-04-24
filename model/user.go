package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId      uint32         `gorm:"column:user_id;type:int UNSIGNED;comment:用户id;primaryKey;not null;" json:"user_id"`                   // 用户id
	UserName    string         `gorm:"column:user_name;type:varchar(20);comment:用户名字;not null;" json:"user_name"`                           // 用户名字
	Password    string         `gorm:"column:password;type:varchar(50);comment:用户密码;not null;" json:"password"`                             // 用户密码
	Mobile      string         `gorm:"column:mobile;type:varchar(11);comment:用户手机号;not null;" json:"mobile"`                                // 用户手机号
	UserBalance int32          `gorm:"column:user_balance;type:int;comment:用户余额;not null;default:0;" json:"user_balance"`                   // 用户余额
	UserPicture string         `gorm:"column:user_picture;type:varchar(200);comment:用户头像;default:NULL;" json:"user_picture"`                // 用户头像
	UserStatus  int8           `gorm:"column:user_status;type:tinyint(1);comment:1正常、2短期封禁 、3永久封禁 ;not null;default:1;" json:"user_status"` // 1正常、2短期封禁 、3永久封禁
	UserScore   int32          `gorm:"column:user_score;type:int;comment:用户积分;not null;default:0;" json:"user_score"`                       // 用户积分
	LastLoginIp string         `gorm:"column:last_login_ip;type:varchar(20);comment:用户最后一次登录的IP地址;default:NULL;" json:"last_login_ip"`      // 用户最后一次登录的IP地址
	Sex         int8           `gorm:"column:sex;type:tinyint(1);comment:1. 男 2. 女 ;default:NULL;" json:"sex"`                              // 1. 男 2. 女
	Age         int32          `gorm:"column:age;type:int;comment:用户年龄;default:NULL;" json:"age"`                                           // 用户年龄
	CreateAt    time.Time      `gorm:"column:create_at;type:datetime;comment:创建时间;default:NULL;" json:"create_at"`                          // 创建时间
	DeleteAt    gorm.DeletedAt `gorm:"column:delete_at;type:datetime;comment:删除时间;default:NULL;" json:"delete_at"`                          // 删除时间
	UpdateAt    time.Time      `gorm:"column:update_at;type:datetime;comment:更新时间;default:NULL;" json:"update_at"`                          // 更新时间
	Account     string         `gorm:"column:account;type:varchar(11);comment:用户账号;default:NULL;" json:"account"`                           // 用户账号
}

func (u *User) TableName() string {
	return "user"
}
