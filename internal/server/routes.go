package server

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/comment"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/content"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/user"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	userRepository := user.NewUserRepository(s.DB)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)

	contentRepository := content.NewContentRepository(s.DB)
	contentService := content.NewContentService(contentRepository)
	contentHandler := content.NewHandler(contentService)

	commentRepository := comment.NewContentRepository(s.DB)
	commentService := comment.NewContentService(commentRepository)
	commentHandler := comment.NewHandler(commentService)

	mux.HandleFunc("POST /user/register", userHandler.Register)
	mux.HandleFunc("POST /user/login", userHandler.Login)
	mux.HandleFunc("GET /user/findall", userHandler.FindAllUsers)
	mux.HandleFunc("GET /user/findone", userHandler.FindUser)
	mux.HandleFunc("PUT /user/update", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /user/delete", userHandler.DeleteUser)

	mux.HandleFunc("POST /content/upload", contentHandler.UploadContent)
	mux.HandleFunc("GET /content/findall", contentHandler.GetAllContent)
	mux.HandleFunc("GET /content/findone", contentHandler.GetOneContent)
	mux.HandleFunc("PUT /content/update", contentHandler.UpdateContent)
	mux.HandleFunc("DELETE /content/delete", contentHandler.DeleteContent)

	mux.HandleFunc("POST /comment/send", commentHandler.SendComment)
	mux.HandleFunc("GET /comment/findall", commentHandler.GetAllComments)
	mux.HandleFunc("PUT /comment/update", commentHandler.UpdateComment)
	mux.HandleFunc("DELETE /comment/delete", commentHandler.DeleteComment)

	return mux
}
