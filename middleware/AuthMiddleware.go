package middleware

import (
	"echo_shop/common"
	"echo_shop/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取authorization header
		tokenString := c.Request().Header.Get("Authorization")
		fmt.Print("请求token", tokenString,"\n")
		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			fmt.Println("111")
			return c.JSON(http.StatusUnauthorized,map[string]interface{}{"code":401,"msg":"权限不足"})
		}
		tokenString = tokenString[7:]
		token,claims,err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			fmt.Println(err)
			fmt.Println("222")
			return c.JSON(http.StatusUnauthorized,map[string]interface{}{"code":401,"msg":"权限不足"})
		}
		//验证通过后获取claim中的userId
		userId := claims.UserId
		DB := common.DB
		var user model.User
		DB.First(&user,userId)

		//用户
		if user.ID == 0 {
			fmt.Println("333")
			return c.JSON(http.StatusUnauthorized,map[string]interface{}{"code":401,"msg":"权限不足"})
		}
		//用户存在，将user的信息写入上下文
		c.Set("user",user)

		return next(c)
	}
}















