package controller

import (
	"echo_shop/common"
	"echo_shop/dto"
	"echo_shop/model"
	"echo_shop/response"
	"echo_shop/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)
//注册
func Register(c echo.Context) error {
	DB := common.DB
	//获取参数
	name := c.FormValue("name")
	telephone := c.FormValue("telephone")
	password := c.FormValue("password")

	//数据验证
	if len(telephone) != 11{
		return response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须是11位")
		//return c.JSON(http.StatusUnprocessableEntity,map[string]interface{}{"code":422,"msg":"手机号必须是11位"})
	}
	if len(password) < 6 {
		return response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不少于6位")
		//return c.JSON(http.StatusUnprocessableEntity,map[string]interface{}{"code":422,"msg":"密码不少于6位"})
	}

	//如果名称未传 / 传空，给一个10位的随机字符串
	if len(name) == 0{
		name = util.RandomString(10)
	}

	fmt.Println(name,password,telephone)

	//查询手机号是否存在
	if isTelephoneExist(DB,telephone){
		return response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		//return c.JSON(http.StatusUnprocessableEntity,map[string]interface{}{"code":422,"msg":"用户已经存在"})
	}
	//创建用户 + 加密用户密码
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return response.Response(c,http.StatusInternalServerError,500,nil,"加密错误")
		//return c.JSON(http.StatusInternalServerError,map[string]interface{}{"code":500,"msg":"加密错误"})
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	return response.Success(c,nil,"注册成功")
	//return c.JSON(200,map[string]interface{}{"msg":"注册成功"})
}

func isTelephoneExist(db *gorm.DB,telephone string) bool{
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}

//登录
func Login(c echo.Context) error {
	DB := common.DB
	//获取参数
	telephone := c.FormValue("telephone")
	password := c.FormValue("password")
	//数据验证
	if len(telephone) != 11{
		return response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须是11位")
		//return c.JSON(http.StatusUnprocessableEntity,map[string]interface{}{"code":422,"msg":"手机号必须是11位"})
	}
	if len(password) < 6 {
		return response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不少于6位")
		//return c.JSON(http.StatusUnprocessableEntity,map[string]interface{}{"code":422,"msg":"密码不少于6位"})
	}
	//判断手机号是否存在
	//if !isTelephoneExist(DB,telephone){
	//		return c.JSON(http.StatusUnprocessableEntity,map[string]interface{}{"code":422,"msg":"用户不存在"})
	//	}
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		return response.Response(c,http.StatusUnprocessableEntity,400,nil,"用户不存在")
		//return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code":400,"msg":"用户不存在"})
	}
	//判断密码是否正确
	//var user model.User
	fmt.Println(user)
	fmt.Println(user.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)) ;err != nil{
		return response.Response(c,http.StatusBadRequest,400,nil,"密码错误")
		//return c.JSON(http.StatusBadRequest,map[string]interface{}{"code":400,"msg":"密码错误"})
	}
	//发放token
	//token := "11"
	//return c.JSON(200,map[string]interface{}{"code":200,"data":map[string]interface{}{"token":token},"msg":"登录成功"})
	token,err := common.ReleaseToken(user)
	if err != nil {
		fmt.Printf("token generate error : %v",err)
		return response.Response(c,http.StatusInternalServerError,500,nil,"系统异常")
		//return c.JSON(http.StatusInternalServerError,map[string]interface{}{"code":500,"msg":"系统异常"})
	} else {
		//return c.JSON(200,map[string]interface{}{
		//		//	"code":200,
		//		//	"data":map[string]interface{}{"token":token},
		//		//	"msg":"登录成功",
		//		//}
		return response.Success(c,map[string]interface{}{"token":token},"登录成功")
	}
}
func Info(c echo.Context) error {
	user := c.Get("user")
	return c.JSON(http.StatusOK,map[string]interface{}{"code":200,"data":map[string]interface{}{"user":dto.ToUserDto(user.(model.User))}})

}