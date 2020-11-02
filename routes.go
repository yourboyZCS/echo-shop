package main

import (
	"echo_shop/controller"
	"echo_shop/middleware"
	"github.com/labstack/echo/v4"
)

func CollectRoute(e *echo.Echo) *echo.Echo{
	e.POST("/api/auth/register", controller.Register)
	e.POST("/api/auth/login", controller.Login)
	//e.Use(middleware.AuthMiddleware)        //错误写法：e.Use(middleware.AuthMiddleware())
	e.GET("/api/auth/info",controller.Info,middleware.AuthMiddleware)
	return e
}
