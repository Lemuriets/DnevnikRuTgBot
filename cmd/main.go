package main

import (
	"log"

	"github.com/Lemuriets/marksBot/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load .env file")
	}

	botApp := app.NewBot("DnevnikRuBot")

}
