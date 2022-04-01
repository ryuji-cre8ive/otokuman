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
}

type SpecificUser struct {
	Id string 
	Name string `json:"name"`
	Password string `json:"password"`
	Flag bool `json:"isCorrectUser"`
}

type DB_Config struct {
	User string `json:"user"`
	Password string `json:"password"`
	DBName string `json:"db_name"`
	Port string `json:"port"`
}

func main () {
  // if err != nil {
  //   panic(err.Error())
  // } else {
  //   println("DB接続成功")
  // }

	// error := db.Create(&User {
	// 	Id: 2,
	// 	Name: "testman",
	// 	Password: "test",
	// }).Error

	// if error != nil {
	// 	println(error)
	// } else {
	// 	println("success to add")
	// }

	e := echo.New()

	e.Static("/", "dist/")
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())

	e.GET("/hello", hello)
	e.POST("/adduser", addUser)
	e.POST("/login", login)

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

	
	dsn := "host=localhost user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " port=" + DBPort + " sslmode=disable TimeZone=Asia/Shanghai"
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
	isCorrectUser := false
	name := param.Name
	password := param.Password
	println("name: " + name + "password: " + password)

	db.First(&user, "name = ?", name)
	if name == user.Name && password == user.Password {
		isCorrectUser = true
	}

	// params, err := json.Marshal(SpecificUser{
	// 	Id: id,
	// 	Name: user.Name,
	// 	Password: password,
	// 	Flag: isCorrectUser,
	// })

	paramsAfterLogin := SpecificUser{
		Id: user.Id,
		Name: user.Name,
		Password: password,
		Flag: isCorrectUser,
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
	fmt.Println("name: " + name + "password: " + password)
	uuidObj, _ := uuid.NewUUID()

	error := db.Create(&User {
		Id: uuidObj.String(),
		Name: name,
		Password: password,
	}).Error

	if error != nil {
		println(error)
	} else {
		println("success to add")
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

	param, err := json.Marshal(DB_Config{
		User: user,
		Password: password,
		DBName: dbname,
		Port: port,
	})

	if err != nil {
		println("Error marshalling")
	} 
	return param	
}

