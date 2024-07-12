package database

import (
	"fmt"
	"os"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err))
	}

	err = db.AutoMigrate(&entity.User{}, &entity.SaveContent{}, &entity.Notification{}, &entity.Message{}, &entity.LikeContent{}, &entity.GroupChat{}, &entity.Follow{}, &entity.Feedback{}, &entity.DirectMessage{}, &entity.Content{}, &entity.Comment{}, &entity.AIMessage{}, &entity.AIConversation{})

	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %s", err))
	}

	fmt.Println("connected to database")

	return db
}
