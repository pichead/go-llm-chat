package config

import (
	"os"
	"github.com/joho/godotenv"
	"log"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		println("Load environment successfully")
	}
}

func getEnv(key string, fallback string) string {


	println("find key : ", key)

	data := os.Getenv(key)
	println("data : ", data)

	if data == "" {
		return fallback
	}
	return data
}

type TEnv struct {
	Port         string
	AppName      string
	AccessToken  string
	RefreshToken string
}

var Env = TEnv{
	Port: getEnv("PORT", "8080"),
}

var App = struct {
	IsProd bool
}{
	IsProd: false,
}
