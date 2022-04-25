package controller

import (
	"github.com/jinzhu/gorm"
	// database "github.com/ryuji-cre8ive/otokuman/database"
	"encoding/json"
	"github.com/ryuji-cre8ive/otokuman/plugins"
	"github.com/ryuji-cre8ive/otokuman/types"
	. "fmt"
)

type User struct {
	gorm.Model
	Password string
	Session string
}

type Profile struct {
	User User
	Image string
	FirstName string
	LastName string
	Address string
	Age int
}

type Article struct {
	gorm.Model
	Title string
	Content string
	genre string
	User User
}

func SqlConnect () (database *gorm.DB, err error)  {
	DBMS := "postgres"
	envData := plugins.LoadEnv()
	var d types.DB_Config
	if err := json.Unmarshal([]byte(envData), &d); err != nil {
		println(err)
	}
	DBUser := d.User
	DBPassword := d.Password
	DBPort := d.Port
	DBName := d.DBName
	DBHost := d.DBHost


	
	dsn := "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " port=" + DBPort + " sslmode=disable TimeZone=Asia/Shanghai"
	Println(dsn)

	return gorm.Open(DBMS, dsn)
}