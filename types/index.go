package types

type User struct {
	Id string 
	Name string `json:"name"`
	Password string `json:"password"`
	Session string
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

type Articles struct {
	Id string `gorm:"" json:"id"`
	Title string `gorm:"" json:"title"`
	Content string `gorm:"" json:"content"`
	Genre string `gorm:"" json:"genre"`
	CreatedAt string `gorm:"" json:"createdAt"`
	CreateUser string `gorm:"foreignkey:Id" json:"createUser"`
}