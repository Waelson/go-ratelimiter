package limiter

type LimiterInterface interface {
	Allow(key string, limit int, blockTime int) bool
}
