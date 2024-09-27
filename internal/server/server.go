package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	port       int
	DB         *gorm.DB
	Redis      *redis.Client
	S3Client   *s3.Client
	BucketName string
}

func NewServer(client *s3.Client, bucketName string) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db := database.New()
	redis := database.NewRedis()

	fmt.Println(db)

	NewServer := &Server{
		port:       port,
		DB:         db,
		S3Client:   client,
		BucketName: bucketName,
		Redis: redis,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}

	return server
}
