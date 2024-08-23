package main

import (
	"context"
	"fmt"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/server"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	client := s3.NewFromConfig(cfg)
	bucketName := "connect-verse-bucket"

	fmt.Println("S3 Client Region:", cfg.Region)

	server := *server.NewServer(client, bucketName)

	err = server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
