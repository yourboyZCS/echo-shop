package main

import (
	"echo_shop/common"
	_ "fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"os"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	//InitConfig()  //从文件中读取配置，未完成，感觉较之前略微麻烦
	//创建数据库 terminal执行
	//连接数据库 建立函数
	err := common.InitMysql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库
	defer common.DB.Close()

	e := echo.New()
	//e.Use(Print)
	e = CollectRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}

//打印请求信息的中间件
//func Print(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		method := c.Request().Method
//		fmt.Println("请求方法：",method)
//		return next(c)
//	}
//}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")      //读取的文件名
	viper.SetConfigType("yml")              //读取的文件类型
	viper.AddConfigPath(workDir + "/config")    //读取的文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

