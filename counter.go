package rcounter

import (
	"strconv"
)

// AddCount ...
func (rc *RCounter) AddCount(key string) (int64, error) {
	return rc.client.Incr(rc.getKey(key)).Result()
}

// AddCountUnique count unique event
func (rc *RCounter) AddCountUnique(key string, event string) (int64, error) {
	return rc.client.SAdd(rc.getKey(key), event).Result()
}

// Count ...
func (rc *RCounter) Count(key string) int64 {
	keyType, err := rc.client.Type(rc.getKey(key)).Result()
	if err != nil {
		return 0
	}
	if keyType == "string" {
		result, _ := rc.client.Get(rc.getKey(key)).Result()
		resultInt, _ := strconv.ParseInt(result, 10, 64)
		return resultInt
	} else if keyType == "set" {
		resultInt, _ := rc.client.SCard(rc.getKey(key)).Result()
		return resultInt
	}
	return 0
}
