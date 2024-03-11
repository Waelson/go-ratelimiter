package limiter

import (
	"context"
	"fmt"
	"github.com/Waelson/go-ratelimit/internal/config"
	"github.com/Waelson/go-ratelimit/internal/storage"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Limiter struct {
	storage storage.StorageInterface
	Config  *config.Configuration
}

func NewLimiter(r storage.StorageInterface) *Limiter {
	return &Limiter{
		storage: r,
	}
}

func (l *Limiter) Allow(key string, limit int, blockTime int) bool {
	ctx := context.Background()
	val, err := l.storage.Get(ctx, key).Result()

	if err == redis.Nil {
		// Chave não existe, então pode prosseguir e criar a chave com limite de 0
		err := l.storage.Set(ctx, key, 1, time.Second).Err()
		if err != nil {
			fmt.Println("Redis set error:", err)
			return false
		}
		return true
	} else if err != nil {
		fmt.Println("Redis get error:", err)
		return false
	}

	// Converte o valor atual para int
	count, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Error converting value:", err)
		return false
	}

	// O IP ou API_KEY esta bloqueado
	if count == -1 {
		fmt.Println(fmt.Printf("IP/API_KEY '%s' is blocked", key))
		return false
	}

	if count > limit {
		// Transforma o block time em time.Duration
		duration := time.Duration(blockTime) * time.Second
		err := l.storage.Set(ctx, key, -1, duration).Err()
		if err != nil {
			fmt.Println("Error setting reached limit", err)
		}
		// Limite atingido
		return false
	}

	// Incrementa a contagem, já que o limite não foi atingido
	err = l.storage.Incr(ctx, key).Err()
	if err != nil {
		fmt.Println("Redis increment error:", err)
		return false
	}
	fmt.Println("Request allowed")
	return true
}
