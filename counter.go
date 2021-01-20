package rcounter

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// Count ...
func (rc *RCounter) Count(key string, args ...time.Duration) int64 {
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
	} else if keyType == "zset" {
		if len(args) > 0 {
			now := time.Now().Unix()
			resultInt, _ := rc.client.ZCount(rc.getKey(key), strconv.FormatInt(now-int64(args[0].Seconds()), 10), strconv.FormatInt(now, 10)).Result()
			return resultInt
		}
		resultInt, _ := rc.client.ZCard(rc.getKey(key)).Result()
		return resultInt
	}
	return 0
}

// AddCount 简单计数
func (rc *RCounter) AddCount(key string) (int64, error) {
	go rc.saveKey(key)
	return rc.client.Incr(rc.getKey(key)).Result()
}

// AddCountUnique count unique event 按照非重复计数
func (rc *RCounter) AddCountUnique(key string, event string) (int64, error) {
	go rc.saveKey(key)
	return rc.client.SAdd(rc.getKey(key), event).Result()
}

// AddCountUniqueInTime count unique event in a period of time 统计一段时间内的非重复计数
func (rc *RCounter) AddCountUniqueInTime(key string, event string, period time.Duration) {
	go rc.saveKey(key)
	rc.client.ZAdd(rc.getKey(key), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: rc.md5(event),
	})
	if rc.AutoDelete {
		// auto delete keys
		go rc.client.ZRemRangeByScore(rc.getKey(key), "0", strconv.FormatInt(time.Now().Unix()-int64(period.Seconds()), 10))
	}
}
