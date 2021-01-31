package context

import (
	"devJoyTelegramBot/pkg/database"
	log "devJoyTelegramBot/pkg/log"
	"devJoyTelegramBot/pkg/util"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var once sync.Once

type BotCtx struct {
	Conn *gorm.DB
}
var instance *BotCtx

func GetCtx()(*BotCtx) {
	once.Do(func() {
		loadEnvs()
		conn := connectDb()
		createTables(conn)
		instance = &BotCtx{Conn: conn}
	})
	return  instance
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		currentWorkingProcessPath, _ := os.Getwd()
		msg :=fmt.Sprintf("%s .env file not found", currentWorkingProcessPath)
		log.Info(msg)
	}
}

func connectDb() *gorm.DB {
	host, user, password, dbname, port := util.GetEnvOrDefault("DB_HOST", "localhost"), os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), util.GetEnvOrDefault("DB_PORT", "5432")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user,password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info("Can't connect to DB")
		panic(err)
	}
	return db
}

func createTables(conn *gorm.DB) {
	conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err := conn.AutoMigrate(&database.Chat{}, &database.Message{})
	if err !=nil {
		panic(err)
	}
}