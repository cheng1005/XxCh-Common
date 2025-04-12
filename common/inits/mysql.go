package inits

import (
	"fmt"
	"github.com/cheng1005/XxCh-Common/common/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() {
	mysqlConfig := global.Nacos.Mysql
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("mysql init success")
}
