package rcounter

import (
	"time"

	"github.com/go-redis/redis"
)

// RCounter ...
type RCounter struct {
	client *redis.Client
	Prefix string
}

// NewRCounter ...
func NewRCounter(Addr string, Password string, DB int) *RCounter {
	rcounter := RCounter{}
	rcounter.Init(Addr, Password, DB)
	rcounter.Prefix = "rcounter_"
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
	newKeys := []string{}
	for _, key := range keys {
		newKeys = append(newKeys, rc.getKey(key))
	}
	rc.client.Del(newKeys...)
}

// Expire expire key
func (rc *RCounter) Expire(key string, expiration time.Duration) {
	rc.client.Expire(rc.getKey(key), expiration)
}

// Exists if key exists
func (rc *RCounter) Exists(key string) bool {
	num, err := rc.client.Exists(rc.getKey(key)).Result()
	if err == nil && num == 1 {
		return true
	}
	return false
}

// TTL get the key expire time
func (rc *RCounter) TTL(key string) (time.Duration, error) {
	return rc.client.TTL(rc.getKey(key)).Result()
}

// DeleteAll delete all the keys user add include the rcount_keys_list
func (rc *RCounter) DeleteAll() {
	list, _ := rc.client.SMembers("rcount_keys_list").Result()
	keys := []string{}
	for _, key := range list {
		keys = append(keys, rc.getKey(key))
	}
	keys = append(keys, "rcount_keys_list")
	rc.client.Del(keys...)
}

func (rc *RCounter) getKey(key string) string {
	return rc.Prefix + key
}

func (rc *RCounter) saveKey(key string) {
	rc.client.SAdd("rcount_keys_list", key)
}
