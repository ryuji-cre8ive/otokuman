package controller

import (
	"github.com/labstack/echo"
	types "github.com/ryuji-cre8ive/otokuman/types"
	plugins "github.com/ryuji-cre8ive/otokuman/plugins"
	"github.com/google/uuid"
	"net/http"
	"time"
	. "fmt"
)

type Users struct {
	Id string 
	Name string `gorm:"" json:"name"`
	Password string `gorm:"" json:"password"`
	Session string
}

func Login (c echo.Context) error {
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	var user Users

	param := new(types.User)
	if err := c.Bind(param); err != nil {
		println(err)
	}
	
	name := param.Name
	password := param.Password
	println("name: " + name + " password: " + password)

	db.First(&user, "name = ?", name)

	//ハッシュ関数にかけて保存してあるパスワードと一致させるためにハッシュ関数にかける
	encryptedPassword := plugins.Encryption(password)

	if name == user.Name {
		if encryptedPassword != user.Password {
			res := plugins.MakeResponse(401, "password is wrong but user is exist")
			return c.JSON(401, res)
		}
	}
	session, _ := uuid.NewUUID()

	cookie := new(http.Cookie)
  cookie.Name = "session"
  cookie.Value = session.String()
  cookie.Expires = time.Now().Add(72 * time.Hour)
  c.SetCookie(cookie)

	user.Session = session.String()
	db.Save(&user)

	paramsAfterLogin := types.SpecificUser{
		Id: user.Id,
		Name: user.Name,
		Password: password,
	}
	Println(paramsAfterLogin)

	return c.JSON(http.StatusOK, paramsAfterLogin)
}


func AddUser (c echo.Context) error {
	param := new(types.User)
	if err := c.Bind(param); err != nil {
		return err
	}
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}


	name := param.Name
	password := param.Password
	Println("name: " + name + " password: " + password)

	
	//ユーザーが既に存在したときの処理
	var user Users
	db.First(&user, "name = ?", name)
	if user.Name != "" {
		res := plugins.MakeResponse(401, "user is already added")
		return c.JSON(http.StatusOK, res)
	}

	uuidObj, _ := uuid.NewUUID()

	encryptedPassword := plugins.Encryption(password)

	error := db.Create(&types.User {
		Id: uuidObj.String(),
		Name: name,
		Password: encryptedPassword,
	}).Error
	Println("encrypted", encryptedPassword)

	if error != nil {
		println(error)
	} else {
		println("success to add user")
	}
	
	return c.JSON(http.StatusOK, param)
}