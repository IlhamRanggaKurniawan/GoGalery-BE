package middleware

import (
	"net/http"
	"strings"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(md ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(md) - 1; i >= 0; i-- {
			next = md[i](next)
		}
		return next
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unprotectedRoutes := []string{
			"/user/register",
			"/user/login",
			"/token",
			"/ws/dm",
		}

		for _, route := range unprotectedRoutes {
			if strings.HasPrefix(r.URL.Path, route) {
				next.ServeHTTP(w, r)
				return
			}
		}

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid token type", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		_, err := utils.ValidateToken(tokenString, "Access Token")

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := []string{
			"https://connect-verse-seven.vercel.app",
			"http://localhost:3000",
			"http://localhost:3001",
			"",
		}

		origin := r.Header.Get("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
