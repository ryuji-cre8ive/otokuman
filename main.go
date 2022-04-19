package main

import (
  "github.com/labstack/echo/middleware"
  _ "github.com/lib/pq"

	"net/url"
	router "github.com/ryuji-cre8ive/otokuman/router"
)

func main () {
	
	e := router.Init()

	url1, err := url.Parse("http://node:3000")
	if err != nil {
		e.Logger.Fatal(err)
	}

	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}

	
	
	// e.Static("/", "dist/")
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	// e.Group("", middleware.Proxy(NoBalancer(url1)))
	g := e.Group("/")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(":1234"))
}