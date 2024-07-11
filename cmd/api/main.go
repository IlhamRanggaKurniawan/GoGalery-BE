package main

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	server := *server.NewServer()

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
