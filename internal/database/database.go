package database

import (
	"fmt"
	"os"
	"time"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := os.Getenv("DB_DSN")

	var db *gorm.DB
	var err error

	for attempts := 0; attempts < 5; attempts++ { // Retry up to 5 times
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break // Successful connection
		}
		fmt.Printf("Error connecting to database: %s. Retrying in 2 seconds...\n", err)
		time.Sleep(2 * time.Second) // Delay before retrying
	}

	if err != nil {
		return nil
	}

	// Auto migrate the entities
	err = db.AutoMigrate(
		&entity.User{},
		&entity.SaveContent{},
		&entity.Notification{},
		&entity.Message{},
		&entity.LikeContent{},
		&entity.GroupChat{},
		&entity.Follow{},
		&entity.Feedback{},
		&entity.DirectMessage{},
		&entity.Content{},
		&entity.Comment{},
		&entity.AIMessage{},
		&entity.AIConversation{},
	)
	if err != nil {
		return nil
	}

	fmt.Println("Connected to database")
	return db
}
