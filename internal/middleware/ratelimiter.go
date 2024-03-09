package middleware

import (
	"github.com/Waelson/go-ratelimit/internal/config"
	"github.com/Waelson/go-ratelimit/internal/limiter"
	"net/http"
)

// NewRateLimiterMiddleware cria e retorna um novo middleware de rate limiting
func NewRateLimiterMiddleware(l *limiter.Limiter, config *config.Configuration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			limit := config.TokenRateLimit
			key := r.Header.Get("API_KEY")

			if key == "" {
				key = r.RemoteAddr
				limit = config.IPRateLimit
			}

			// Verifica primeiro o limite baseado no token, se disponível
			if key != "" && !l.Allow(key, limit, config.BlockDuration) {
				httpError(w, "you have reached the maximum number of requests allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			// Se nenhum limite foi atingido, processa a requisição
			next.ServeHTTP(w, r)
		})
	}
}

// httpError envia uma resposta de erro HTTP com uma mensagem específica e um código de status
func httpError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(message))
}
