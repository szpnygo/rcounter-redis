package rcounter

import (
	"strconv"
)

// AddCount ...
func (rc *RCounter) AddCount(key string) (int64, error) {
	return rc.client.Incr(key).Result()
}

// AddCountUnique count unique event
func (rc *RCounter) AddCountUnique(key string, event string) (int64, error) {
	return rc.client.SAdd(key, event).Result()
}

// Count ...
func (rc *RCounter) Count(key string) int64 {
	keyType, err := rc.client.Type(key).Result()
	if err != nil {
		return 0
	}
	if keyType == "string" {
		result, _ := rc.client.Get(key).Result()
		resultInt, _ := strconv.ParseInt(result, 10, 64)
		return resultInt
	} else if keyType == "set" {
		resultInt, _ := rc.client.SCard(key).Result()
		return resultInt
	}
	return 0
}
