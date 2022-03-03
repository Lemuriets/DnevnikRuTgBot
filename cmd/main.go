package main

import (
	"fmt"
	"log"

	// "github.com/Lemuriets/marksBot/internal/app"
	"github.com/Lemuriets/marksBot/pkg/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load .env file")
	}

	content := http.Login("https://login.dnevnik.ru/login/", "yegormadyarov", "egormadyarov9")

	fmt.Println(content)
	// botApp := app.NewBot("DnevnikRuBot")
	// botApp.Start()
}
