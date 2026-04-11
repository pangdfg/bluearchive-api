package config

import (
	"strings"
	"time"

	env "bluearchive-go/internal/env"
)

type Config struct {
	Server    ServerConfig
	Upstream  UpstreamConfig
	Cache     CacheConfig
	Auth      AuthConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type UpstreamConfig struct {
	BaseURL    string
	Timeout    time.Duration
	MaxRetries int
}

type CacheConfig struct { 
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	RedisPrefix   string
	RedisPoolSize int
}

type AuthConfig struct {
	APIKeys []string // empty = auth disabled
}

type RateLimitConfig struct {
	Rate    float64 // tokens per second
	Burst   float64 // max burst
	Enabled bool
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         env.GetEnv("PORT", "8080"),
			ReadTimeout:  env.GetDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: env.GetDuration("SERVER_WRITE_TIMEOUT", 30*time.Second),
			IdleTimeout:  env.GetDuration("SERVER_IDLE_TIMEOUT", 60*time.Second),
		},
		Upstream: UpstreamConfig{
			BaseURL:    env.GetEnv("UPSTREAM_BASE_URL", "https://api.ennead.cc/buruaka"),
			Timeout:    env.GetDuration("UPSTREAM_TIMEOUT", 15*time.Second),
			MaxRetries: env.GetInt("UPSTREAM_MAX_RETRIES", 3),
		},
		Cache: CacheConfig{
			RedisAddr:       env.GetEnv("REDIS_ADDR", "localhost:6379"),
			RedisPassword:   env.GetEnv("REDIS_PASSWORD", ""),
			RedisDB:         env.GetInt("REDIS_DB", 0),
			RedisPrefix:     env.GetEnv("REDIS_KEY_PREFIX", "ba:"),
			RedisPoolSize:   env.GetInt("REDIS_POOL_SIZE", 10),
		},
		Auth: AuthConfig{
			APIKeys: parseList(env.GetEnv("API_KEYS", "")),
		},
		RateLimit: RateLimitConfig{
			Rate:    env.GetFloat("RATE_LIMIT_RPS", 10),
			Burst:   env.GetFloat("RATE_LIMIT_BURST", 30),
			Enabled: env.GetBool("RATE_LIMIT_ENABLED", true),
		},
	}
}

func parseList(s string) []string {
	if s == "" {
		return nil
	}
	var out []string
	for _, part := range strings.Split(s, ",") {
		if t := strings.TrimSpace(part); t != "" {
			out = append(out, t)
		}
	}
	return out
}