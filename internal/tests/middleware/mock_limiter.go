package middleware_test

import (
	"fmt"
)

// MockLimiter é uma implementação mock do Limiter para testes
type MockLimiter struct {
	AllowFunc func(key string, limit int, blockDuration int) bool
}

func (m *MockLimiter) Allow(key string, limit int, blockDuration int) bool {
	if m.AllowFunc != nil {
		return m.AllowFunc(key, limit, blockDuration)
	}
	fmt.Println("MockLimiter Allow method not implemented")
	return false
}
