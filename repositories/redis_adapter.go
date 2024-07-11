package repositories

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisRepository struct {
	db *redis.Client
}

func NewRedisRepository(db *redis.Client) RedisRepository {
	return &redisRepository{db: db}
}

func (r *redisRepository) Save(key string, otp string, expiration time.Duration) error {
	err := r.db.Set(key, otp, expiration).Err()
	if err != nil {
		fmt.Println("error repo:", err)
		return err
	}
	return nil
}

func (r *redisRepository) Get(key string) (string, error) {
	otp, err := r.db.Get(key).Result()
	if err != redis.Nil {
		return "", fmt.Errorf("OTP not found or expired")
	} else if err != nil {
		return "", err
	}
	return otp, nil
}
