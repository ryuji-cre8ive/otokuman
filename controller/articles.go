package controller

import (
	"github.com/labstack/echo"
	types "github.com/ryuji-cre8ive/otokuman/types"
	"github.com/google/uuid"
	"fmt"
	plugins "github.com/ryuji-cre8ive/otokuman/plugins"
	"net/http"
)

func GetArticles (c echo.Context) error {
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	articles := []types.Articles{}

	results := db.Find(&articles)
	fmt.Println("result")
	fmt.Println(results)

	return c.JSON(http.StatusOK, results)
}

func AddArticle (c echo.Context) error {
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	param := new(types.Articles)
	if err := c.Bind(param); err != nil {
		println(err)
	}

	userData := GetUserDataWithCookie(c)
	uuidObj, _ := uuid.NewUUID()
	title := param.Title
	content := param.Content
	genre := param.Genre
	fmt.Println(userData.Id)
	
	createdAt := plugins.GetDate()

	error := db.Create(&types.Articles {
		Id: uuidObj.String(),
		Title: title,
		Content: content,
		Genre: genre,
		CreatedAt: createdAt,
		CreateUser: userData.Id,
	}).Error

	if error != nil {
		println(error)
	} else {
		println("success to add article")
	}

	res := plugins.MakeResponse(200, "your article is successfuly added")
	
	return c.JSON(http.StatusOK, res)
}

func GetArticlesWrittenByMe (c echo.Context) error {
	cookie, err := c.Cookie("session")
		if err != nil {
			return err
	}
	db, error := SqlConnect()
	if error != nil {
		return err
	}
	articles := []types.Articles{}
	var user types.Users
	db.First(&user, "session = ?", cookie.Value)

	fmt.Println("userId" + user.Id)

	db.Where("create_user = ?", user.Id).Find(&articles)

	fmt.Println(articles)

	return c.JSON(http.StatusOK, articles)

}