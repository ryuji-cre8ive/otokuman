package controller

import (
	"github.com/labstack/echo"
	types "github.com/ryuji-cre8ive/otokuman/types"
	"github.com/google/uuid"
	"fmt"
	plugins "github.com/ryuji-cre8ive/otokuman/plugins"
	"net/http"
)

func AddArticle (c echo.Context) error {
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	param := new(types.Articles)
	if err := c.Bind(param); err != nil {
		println(err)
	}
	uuidObj, _ := uuid.NewUUID()
	title := param.Title
	content := param.Content
	genre := param.Genre
	createUser := param.CreateUser
	
	createdAt := plugins.GetDate()
	fmt.Println("createUser: ")
	fmt.Println(createUser)

	error := db.Create(&types.Articles {
		Id: uuidObj.String(),
		Title: title,
		Content: content,
		Genre: genre,
		CreatedAt: createdAt,
		CreateUser: createUser,
	}).Error

	if error != nil {
		println(error)
	} else {
		println("success to add article")
	}

	res := plugins.MakeResponse(200, "your article is successfuly added")
	
	return c.JSON(http.StatusOK, res)
}