package rcounter

import (
	"time"

	"github.com/go-redis/redis"
)

// RCounter ...
type RCounter struct {
	client *redis.Client
}

// NewRCounter ...
func NewRCounter(Addr string, Password string, DB int) *RCounter {
	rcounter := RCounter{}
	rcounter.Init(Addr, Password, DB)
	return &rcounter
}

// Init ...
func (rc *RCounter) Init(Addr string, Password string, DB int) error {
	rc.client = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
	_, err := rc.client.Ping().Result()
	return err
}

// GetClient ...
func (rc *RCounter) GetClient() *redis.Client {
	return rc.client
}

// Close ...
func (rc *RCounter) Close() {
	rc.client.Close()
}

// Del delete the event
func (rc *RCounter) Del(keys ...string) {
	rc.client.Del(keys...)
}

// Expire expire key
func (rc *RCounter) Expire(key string, expiration time.Duration) {
	rc.client.Expire(key, expiration)
}

// Exists if key exists
func (rc *RCounter) Exists(key string) bool {
	num, err := rc.client.Exists(key).Result()
	if err == nil && num == 1 {
		return true
	}
	return false
}

// TTL get the key expire time
func (rc *RCounter) TTL(key string) (time.Duration, error) {
	return rc.client.TTL(key).Result()
}
