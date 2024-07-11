package repositories

import "time"

type RedisRepository interface {
	Save(string, string, time.Duration) error
	Get(string) (string, error)
}
