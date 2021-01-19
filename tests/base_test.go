package tests_test

import (
	"fmt"
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	key := "new_count"
	fmt.Print("new count key : " + key + " ")
	fmt.Println(rCounter.AddCount(key))
	fmt.Println("is new_count exists ", rCounter.Exists(key))
	fmt.Println("set key expire")
	rCounter.Expire(key, 60*time.Second)
	time.Sleep(2 * time.Second)
	fmt.Print("get the key ttl ")
	fmt.Println(rCounter.TTL(key))
	fmt.Println("del the key")
	rCounter.Del(key)
	fmt.Println("is new_count exists ", rCounter.Exists(key))
}
