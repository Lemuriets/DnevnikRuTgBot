package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Hello I'm Dnevnik.ru bot who can to say marks or shedule from dnevnik.ru for you :)")
	bot.Send(newMsg)
}
