package middleware_test

import (
	"github.com/Waelson/go-ratelimit/internal/config"
	"github.com/Waelson/go-ratelimit/internal/middleware"
	middleware_test "github.com/Waelson/go-ratelimit/internal/tests/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRateLimiterMiddleware(t *testing.T) {
	mockLimiter := &middleware_test.MockLimiter{
		AllowFunc: func(key string, limit int, blockDuration int) bool {
			// Simula a permissão ou negação baseada na chave ou limites
			return key != "deny"
		},
	}

	config := &config.Configuration{
		IPRateLimit:        5,
		TokenRateLimit:     100,
		TokenBlockDuration: 300,
		IPBlockDuration:    300,
	}

	middleware := middleware.NewRateLimiterMiddleware(mockLimiter, config)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Teste com API_KEY que passa a limitação de taxa
	req := httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Add("API_KEY", "allow")
	w := httptest.NewRecorder()
	middleware(testHandler).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Teste com API_KEY que não passa a limitação de taxa
	req = httptest.NewRequest("GET", "http://example.com", nil)
	req.Header.Add("API_KEY", "deny")
	w = httptest.NewRecorder()
	middleware(testHandler).ServeHTTP(w, req)

	if w.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status %d, got %d", http.StatusTooManyRequests, w.Code)
	}

	// Adicione mais cenários conforme necessário
}
