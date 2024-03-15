package middleware

import (
	"fmt"
	"github.com/Waelson/go-ratelimit/internal/config"
	"github.com/Waelson/go-ratelimit/internal/limiter"
	"net"
	"net/http"
	"strings"
)

// NewRateLimiterMiddleware cria e retorna um novo middleware de rate limiting
func NewRateLimiterMiddleware(l limiter.LimiterInterface, config *config.Configuration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			limit := config.TokenRateLimit
			blockDuration := config.TokenBlockDuration
			key := r.Header.Get("API_KEY")
			fmt.Println("API_KEY: ", key)
			if key == "" {
				key = ExtractIP(r.RemoteAddr)
				limit = config.IPRateLimit
				blockDuration = config.IPBlockDuration
			}

			// Verifica primeiro o limite baseado no token, se disponível
			if key != "" && !l.Allow(key, limit, blockDuration) {
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

func ExtractIP(addr string) string {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		// Se houver um erro, pode ser um endereço IP sem porta.
		// Tente tratar como um endereço IP puro.
		if strings.Contains(addr, "[") && strings.Contains(addr, "]") {
			// Endereço IPv6 com colchetes mas sem porta.
			return addr
		}
		// Assume que o erro foi devido à falta de porta e retorna o addr diretamente.
		return addr
	}

	if strings.Contains(ip, ":") && !strings.Contains(ip, "[") {
		// É um IPv6 sem colchetes, adicione-os.
		return "[" + ip + "]"
	}

	// Retorna o IP diretamente (IPv4 ou IPv6 já formatado).
	return ip
}
