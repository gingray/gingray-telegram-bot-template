package database

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ChatId string `json:"chat_id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Messages []*Message `json:"messages"`
}

type Message struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ChatID uuid.UUID
	Chat *Chat `json:"chat"`
	Msg string `json:"msg"`
}


