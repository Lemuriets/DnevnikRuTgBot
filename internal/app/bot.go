package app

import (
	"fmt"
	"os"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Bot     *tgBotApi.BotAPI
	Webhook *tgBotApi.WebhookConfig
}

func NewBot(botName string) *App {
	bot, err := tgBotApi.NewBotAPI(os.Getenv("token"))
	bot.Debug = true // не забудь убрать потом

	if err != nil {
		panic(err)
	}

	fmt.Printf("The bot started as %s", bot.Self.UserName)

	return &App{
		Bot: bot,
	}
}
