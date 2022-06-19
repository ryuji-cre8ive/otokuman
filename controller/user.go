package controller

import (
	"github.com/labstack/echo"
	types "github.com/ryuji-cre8ive/otokuman/types"
	plugins "github.com/ryuji-cre8ive/otokuman/plugins"
	"github.com/google/uuid"
	"net/http"
	"time"
	. "fmt"
	"golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"encoding/base64"
	"strings"
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



func Logout (c echo.Context) error {
	// dbからセッション情報を削除する
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	var userBefore Users

	param := new(types.User)
	if err := c.Bind(param); err != nil {
		println(err)
	}

	userBefore.Id = param.Id
	userAfter := userBefore

	db.First(&userAfter)
	userAfter.Session = ""
	db.Model(&userBefore).Update(&userAfter)
	db.Save(&userAfter)

	// クッキーに保存されているセッション情報を削除する
	cookie, err := c.Cookie("session")
		if err != nil {
			return err
	}
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "you logout")
}



//メールを送信するミドルウェア
func Sendmail (c echo.Context) error {
	type Email struct{
		Mail string
	}

	param := new(Email)
	if err := c.Bind(param); err != nil {
		println(err)
	}

	mail := param.Mail
	
	config := oauth2.Config {
		ClientID: "474021763751-db0s6kjc0u3arg67t6k4jo6m38pv923a.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint: google.Endpoint,
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		Scopes:       []string{"https://mail.google.com/"},
	}
	expiry, _ := time.Parse("2006-01-02", "2017-07-11")

	token := oauth2.Token{
		AccessToken: "https://oauth2.googleapis.com/token",
		TokenType: "Bearer",
		RefreshToken: "https://www.googleapis.com/oauth2/v1/certs",
		Expiry: expiry,
	}

	client := config.Client(oauth2.NoContext, &token)

	srv, err := gmail.New(client)
	if err != nil {
		Printf("Unable to retrieve gmail Client %v", err)
	}

	template:= []byte("From: 'me'\r\n" + 
		"reply-to: otokuman.jp@gmail.com\r\n" + 
		"To: " + mail +"\r\n" +
		"Subject: Test\r\n" + 
		"\r\n" + "TestBody")
	
		var message gmail.Message
		message.Raw = base64.StdEncoding.EncodeToString(template)
		message.Raw = strings.Replace(message.Raw, "/", "_", -1)
		message.Raw = strings.Replace(message.Raw, "+", "-", -1)
		message.Raw = strings.Replace(message.Raw, "=", "", -1)

		_, err = srv.Users.Messages.Send("me", &message).Do()
		if err != nil {
			Printf("%v", err)
		}

		return c.JSON(http.StatusOK, "the system sent email for you")
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
	
	return c.JSON(http.StatusOK, "you are registered the system")
}

func GetUserDataWithCookieHandler (c echo.Context) error {
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	cookie, err := c.Cookie("session")
	if err != nil {
		return err
	}


	var user Users

	session := cookie.Value

	db.Where(&user, "session = ?", session).First(&user)
	gotSessionFromDB := user.Session

	if gotSessionFromDB != session {
		Println("gotDBSession: " + gotSessionFromDB)
		Println("session: " + session)
		return c.String(401, "your cookie is already exipired...")
	}

	paramsAfterLoginWithCookie := types.SpecificUser{
		Id: user.Id,
		Name: user.Name,
		Password: user.Password,
	}
	Println(paramsAfterLoginWithCookie)
	return c.JSON(http.StatusOK, paramsAfterLoginWithCookie)
}

func GetUserDataWithCookie(c echo.Context) (types.Users) {
	db, err := SqlConnect()
	if err != nil {
		println(err)
	}

	cookie, err := c.Cookie("session")
	if err != nil {
		panic(err)
	}


	var user Users

	session := cookie.Value

	db.Where(&user, "session = ?", session).First(&user)
	gotSessionFromDB := user.Session

	if gotSessionFromDB != session {
		Println("gotDBSession: " + gotSessionFromDB)
		Println("session: " + session)
	}

	paramsAfterLoginWithCookie := types.Users{
		Id: user.Id,
		Name: user.Name,
		Password: user.Password,
	}

	return paramsAfterLoginWithCookie
}