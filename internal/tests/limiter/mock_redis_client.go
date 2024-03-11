package limiter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

// MockRedisClient é uma implementação mock da RedisClientInterface.
type MockRedisClient struct {
	data sync.Map // Usado para simular o armazenamento do Redis.
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{}
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	cmd := redis.NewStringCmd(ctx)

	if val, ok := m.data.Load(key); ok {
		cmd.SetVal(val.(string))
	} else {
		cmd.SetErr(redis.Nil)
	}

	return cmd
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(ctx)
	m.data.Store(key, value)
	return cmd
}

func (m *MockRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	cmd := redis.NewIntCmd(ctx)

	val, _ := m.data.LoadOrStore(key, "0")
	current, _ := strconv.Atoi(val.(string))
	current++
	m.data.Store(key, strconv.Itoa(current))
	cmd.SetVal(int64(current))

	return cmd
}
