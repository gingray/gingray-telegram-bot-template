package main

import (
	"devJoyTelegramBot/pkg/context"
	"devJoyTelegramBot/pkg/core"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	logCustom "devJoyTelegramBot/pkg/log"
	"log"
	"net/http"
	"os"
)

func main() {
	context.GetCtx()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(os.Getenv("WEBHOOK_URL")+bot.Token))

	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	listenAddr :=fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT"))
	go http.ListenAndServe(listenAddr, nil)
	//go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for payload := range updates {
		data,_ := json.Marshal(payload)
		core.ProcessWebhook(bot, payload)
		logCustom.InfoJson("Webhook Payload", string(data))
	}
}