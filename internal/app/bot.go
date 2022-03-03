package app

import (
	"fmt"
	"os"

	"github.com/Lemuriets/marksBot/internal/app/commands"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Bot            *tgBotApi.BotAPI
	Webhook        *tgBotApi.WebhookConfig
	UpdatesTimeout int
}

func NewBot(botName string) *App {
	bot, err := tgBotApi.NewBotAPI(os.Getenv("token"))
	bot.Debug = true // не забудь убрать потом

	if err != nil {
		panic(err)
	}

	fmt.Printf("The bot started as %s\n", bot.Self.UserName)

	return &App{
		Bot:            bot,
		UpdatesTimeout: 60,
	}
}

func (botApp *App) Start() {
	u := tgBotApi.NewUpdate(0)
	u.Timeout = botApp.UpdatesTimeout

	updates := botApp.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.Text {
			case "/start":
				commands.Start(botApp.Bot, update.Message)
			default:
				botApp.Bot.Send(tgBotApi.NewMessage(
					update.Message.Chat.ID,
					"Используй /start, чтобы посмотреть список доступных команд",
				))
			}
		}
	}
}
