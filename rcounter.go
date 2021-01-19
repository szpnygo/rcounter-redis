package rcounter

import (
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
