package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// StorageInterface define os métodos usados do cliente Redis.
type StorageInterface interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
}
