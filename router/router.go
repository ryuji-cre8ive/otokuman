package router

import (
	"github.com/labstack/echo"
	con "github.com/ryuji-cre8ive/otokuman/controller"
	midd "github.com/ryuji-cre8ive/otokuman/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("api")

	// api.Use()
	api.GET("/checkCookie",midd.CheckCookieForRouter)
	api.POST("/adduser", con.AddUser)
	api.POST("/login", con.Login)
	api.POST("/addArticle", con.AddArticle, midd.ReadCookieMiddleware)
	api.GET("/getUserDataWithCookie", con.GetUserDataWithCookieHandler)
	api.GET("/getArticles", con.GetArticles)

	return e
}
