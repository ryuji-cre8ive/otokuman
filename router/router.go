package router

import (
	"github.com/labstack/echo"
	con "github.com/ryuji-cre8ive/otokuman/controller"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("api") 
	{
		api.POST("/adduser", con.AddUser)
		api.POST("/login", con.Login)
		api.POST("/addArticle", con.AddArticle)
	}

	return e
}
