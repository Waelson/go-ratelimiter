package main

import (
	"github.com/Waelson/go-ratelimit/internal/config"
	limiter2 "github.com/Waelson/go-ratelimit/internal/limiter"
	"github.com/Waelson/go-ratelimit/internal/middleware"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func main() {
	envLoader := config.GodotenvLoader{}
	configuration, err := config.LoadConfig(&envLoader)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     configuration.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	limiter := limiter2.NewLimiter(redisClient)
	limiterMiddleware := middleware.NewRateLimiterMiddleware(limiter, configuration)

	http.Handle("/", limiterMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome!"))
	})))

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
