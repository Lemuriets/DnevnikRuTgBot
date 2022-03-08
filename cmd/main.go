package main

import (
	"fmt"
	"log"
	"os"

	// "os"

	// "github.com/Lemuriets/marksBot/internal/app"
	exHttp "github.com/Lemuriets/marksBot/pkg/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load .env file")
	}

	h := exHttp.NewHttp()
	resp := h.LoginPost("https://dnevnik.ru", os.Getenv("login"), os.Getenv("password"))

	fmt.Println(resp.Cookies())
	// botApp := app.NewBot("DnevnikRuBot")
	// botApp.Start()
}
