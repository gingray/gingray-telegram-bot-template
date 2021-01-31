package core

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func SendTGTextMessage(bot *tgbotapi.BotAPI, chatID int64, msg string) {
	tgMsg := tgbotapi.NewMessage(chatID, msg)
	_,_ = bot.Send(tgMsg)
}
