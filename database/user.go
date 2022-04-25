package database

import (
	"github.com/jinzhu/gorm"
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