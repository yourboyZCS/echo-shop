package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
返回：
{
	code :200
	data :xxx
	msg : xxx
}
*/

func Response(c echo.Context,httpStatus int,code int,data map[string]interface{},msg string) error {
	return c.JSON(httpStatus,map[string]interface{}{"code":code,"data":data,"msg":msg})
}

func Success(c echo.Context,data map[string]interface{},msg string) error{
	return Response(c,http.StatusOK,200,data,msg)
}

func Fail(c echo.Context,data map[string]interface{},msg string) error{
	return Response(c,http.StatusBadRequest,400,data,msg)
}
