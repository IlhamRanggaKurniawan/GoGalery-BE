package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database"
	"gorm.io/gorm"
)

type Server struct {
	port int
	DB *gorm.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	fmt.Println("port:", port)

	db := database.New()

	NewServer := &Server{
		port: port,
		DB: db,
	}

	server := &http.Server{
		Addr:  fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}

	return server
}