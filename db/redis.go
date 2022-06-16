package db

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/youlance/auth/pkg/config"
)

type DB struct {
	store *redis.Client
}

func New(config config.Config) (*DB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if pong == "" && err != nil {
		return nil, err
	}

	return &DB{store: client}, nil
}

func (s *DB) Set(key string, val string, exp time.Duration) error {
	err := s.store.Set(key, val, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *DB) Get(key string) (string, error) {
	val, err := s.store.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
