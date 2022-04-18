package main

import (
	"net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	// "strconv"
	"encoding/json"
  _ "github.com/lib/pq"
	"github.com/google/uuid"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"time"
)

type User struct {
	Id string 
	Name string `json:"name"`
	Password string `json:"password"`
}

type Users struct {
	Id string 
	Name string `gorm:"" json:"name"`
	Password string `gorm:"" json:"password"`
	Session string
}

type SpecificUser struct {
	Id string 
	Name string `json:"name"`
	Password string `json:"password"`
}

type DB_Config struct {
	User string `json:"user"`
	Password string `json:"password"`
	DBName string `json:"db_name"`
	Port string `json:"port"`
	DBHost string `json:"db_host"`
}

type responseParams struct{
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

type Articles struct {
	Id string `gorm:"" json:"id"`
	Title string `gorm:"" json:"title"`
	Content string `gorm:"" json:"content"`
	Genre string `gorm:"" json:"genre"`
	CreatedAt string `gorm:"" json:"createdAt"`
	CreateUser Users `gorm:"foreignkey:Id" json:"createUser"`
}

func main () {
	e := echo.New()

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
	
	api := e.Group("/api")
	api.GET("/hello", hello)
	api.POST("/adduser", addUser)
	api.POST("/login", login)
	api.POST("/addArticle", addArticle)

	g := e.Group("/")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(":1234"))

	
}

func hello (c echo.Context) error {
	return c.String(http.StatusOK, "hello World!!")
}

func sqlConnect () (database *gorm.DB, err error) {
	DBMS := "postgres"
	envData := loadEnv()
	var d DB_Config
	if err := json.Unmarshal([]byte(envData), &d); err != nil {
		println(err)
	}
	DBUser := d.User
	DBPassword := d.Password
	DBPort := d.Port
	DBName := d.DBName
	DBHost := d.DBHost

	
	
	dsn := "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " port=" + DBPort + " sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)

	return gorm.Open(DBMS, dsn)
}

func login (c echo.Context) error {
	db, err := sqlConnect()
	if err != nil {
		println(err)
	}

	var user Users

	param := new(User)
	if err := c.Bind(param); err != nil {
		println(err)
	}
	
	name := param.Name
	password := param.Password
	println("name: " + name + " password: " + password)

	db.First(&user, "name = ?", name)

	//ハッシュ関数にかけて保存してあるパスワードと一致させるためにハッシュ関数にかける
	encryptedPassword := encryption(password)

	if name == user.Name {
		if encryptedPassword != user.Password {
			res := makeResponse(401, "password is wrong but user is exist")
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

	paramsAfterLogin := SpecificUser{
		Id: user.Id,
		Name: user.Name,
		Password: password,
	}
	fmt.Println(paramsAfterLogin)

	return c.JSON(http.StatusOK, paramsAfterLogin)
}



func addUser (c echo.Context) error {
	param := new(User)
	if err := c.Bind(param); err != nil {
		return err
	}
	db, err := sqlConnect()
	if err != nil {
		println(err)
	}

	name := param.Name
	password := param.Password
	fmt.Println("name: " + name + " password: " + password)

	
	//ユーザーが既に存在したときの処理
	var user Users
	db.First(&user, "name = ?", name)
	if user.Name != "" {
		res := makeResponse(401, "user is already added")
		return c.JSON(http.StatusOK, res)
	}

	uuidObj, _ := uuid.NewUUID()

	encryptedPassword := encryption(password)

	error := db.Create(&User {
		Id: uuidObj.String(),
		Name: name,
		Password: encryptedPassword,
	}).Error
	fmt.Println("encrypted", encryptedPassword)

	if error != nil {
		println(error)
	} else {
		println("success to add user")
	}
	
	
	return c.JSON(http.StatusOK, param)
}

func loadEnv () []byte {
	err := godotenv.Load(".env")
	if err != nil {
		println("Error loading environment", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")

	param, err := json.Marshal(DB_Config{
		User: user,
		Password: password,
		DBName: dbname,
		Port: port,
		DBHost: dbhost,
	})

	if err != nil {
		println("Error marshalling")
	} 
	return param	
}



func encryption (word string) string {
	b := []byte(word)
	sha256 := sha256.Sum256(b)
	hashed := hex.EncodeToString(sha256[:])
	return hashed
}

func makeResponse (code int, message string) responseParams {
	response := responseParams{
		StatusCode: code,
		Message: message,
	}
	return response
}

func addArticle (c echo.Context) error {
	db, err := sqlConnect()
	if err != nil {
		println(err)
	}

	param := new(Articles)
	if err := c.Bind(param); err != nil {
		println(err)
	}
	uuidObj, _ := uuid.NewUUID()
	title := param.Title
	content := param.Content
	genre := param.Genre
	createUser := param.CreateUser
	
	createdAt := getDate()
	fmt.Println("createUser: ")
	fmt.Println(createUser)

	error := db.Create(&Articles {
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

	res := makeResponse(200, "your article is successfuly added")
	
	return c.JSON(http.StatusOK, res)
}

func getDate () string{
	jst, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(jst).Format("2006-01-02-15:04:05")
}

