package database

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func FindOrCreateChat(conn *gorm.DB, chatID string, chatType string, title string) *Chat {
	var chat *Chat
	result :=conn.Where("chat_id = ?", chatID).First(&chat)
	if result.Error == nil {
		return chat
	}
	chat = &Chat{ChatId: chatID, Type: chatType, Title: title}
	result = conn.Create(chat)

	return chat
}

func CreateMessage(conn *gorm.DB, chatID uuid.UUID, msg string) {
	conn.Create(&Message{ChatID: chatID, Msg: msg})
}