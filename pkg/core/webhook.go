package core

import (
	"devJoyTelegramBot/pkg/context"
	"devJoyTelegramBot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func ProcessWebhook(bot *tgbotapi.BotAPI, payload tgbotapi.Update) {
	ctx := context.GetCtx()
	if !isValidPayload(payload) {
		return
	}

	chatID :=strconv.FormatInt(payload.ChannelPost.Chat.ID,10)
	chat := database.FindOrCreateChat(ctx.Conn, chatID, payload.ChannelPost.Chat.Type, payload.ChannelPost.Chat.Title)
	msg := payload.ChannelPost.Text
	database.CreateMessage(ctx.Conn, chat.ID, msg)
	SendTGTextMessage(bot,payload.ChannelPost.Chat.ID,msg)
	tgMsg := tgbotapi.NewMessage(payload.ChannelPost.Chat.ID, msg)
	_,_ = bot.Send(tgMsg)
}

func isValidPayload(payload tgbotapi.Update) bool {
	return payload.ChannelPost !=nil && payload.ChannelPost.Chat !=nil
}
