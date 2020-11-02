package common

import (
	"echo_shop/model"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)
func InitMysql() (err error){
	//不能用:= , 全局
	DB, err = gorm.Open("mysql","root:123456@(127.0.0.1)/echo_pro1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	//模型绑定
	DB.AutoMigrate(&model.User{})
	return DB.DB().Ping()
}