package limiter_test

import (
	"context"
	"github.com/Waelson/go-ratelimit/internal/limiter"
	limiter2 "github.com/Waelson/go-ratelimit/internal/tests/limiter"
	"testing"
)

func TestLimiter_Allow(t *testing.T) {
	mockRedis := limiter2.NewMockRedisClient()
	l := limiter.NewLimiter(mockRedis)

	tests := []struct {
		name      string
		key       string
		limit     int
		blockTime int
		want      bool
		prepMock  func()
	}{
		{
			name:      "NewKey_Allowed",
			key:       "new_key",
			limit:     10,
			blockTime: 60,
			want:      true,
			prepMock:  func() {},
		},
		{
			name:      "ExistingKey_Allowed",
			key:       "existing_key",
			limit:     1,
			blockTime: 60,
			want:      true,
			prepMock: func() {
				mockRedis.Set(context.Background(), "existing_key", "1", 0)
			},
		},
		// Adicione mais cenários conforme necessário.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepMock() // Prepara o mock conforme necessário.
			if got := l.Allow(tt.key, tt.limit, tt.blockTime); got != tt.want {
				t.Errorf("Limiter.Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}
