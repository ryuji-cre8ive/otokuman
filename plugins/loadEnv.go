package plugins

import (
	"os"
	"github.com/joho/godotenv"
	"encoding/json"
	types "github.com/ryuji-cre8ive/otokuman/types"
)

func LoadEnv () []byte {
	err := godotenv.Load(".env")
	if err != nil {
		println("Error loading environment", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")

	param, err := json.Marshal(types.DB_Config{
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