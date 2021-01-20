package tests_test

import (
	"fmt"
	"testing"
	"time"
)

func TestAddCount(t *testing.T) {
	num, err := rCounter.AddCount("count")
	fmt.Println("count", num, err)
	fmt.Println("count", rCounter.Count("count"))
}

func TestAddCountUnique(t *testing.T) {
	num, err := rCounter.AddCountUnique("unique_count", "user_id_1")
	fmt.Println("add unique 1 :", num, err)
	num, err = rCounter.AddCountUnique("unique_count", "user_id_2")
	fmt.Println("add unique 2 :", num, err)
	num, err = rCounter.AddCountUnique("unique_count", "user_id_3")
	fmt.Println("add unique 3 :", num, err)
	fmt.Println("unique count", rCounter.Count("unique_count"))
}

func TestAddCountUniqueInTime(t *testing.T) {
	rCounter.AutoDelete = false
	key := "unique_in_time"
	fmt.Println("add the event user_1_click and sleep 1 second")
	rCounter.AddCountUniqueInTime(key, "user_1_click", 5*time.Second)
	time.Sleep(1 * time.Second)

	fmt.Println("add the event user_1_click again and user_2_click")
	rCounter.AddCountUniqueInTime(key, "user_1_click", 5*time.Second)
	rCounter.AddCountUniqueInTime(key, "user_2_click", 5*time.Second)
	fmt.Println("now count ", rCounter.Count(key, 2*time.Second), " and sleep 2 second")

	time.Sleep(3 * time.Second)
	fmt.Println("end count ", rCounter.Count(key, 2*time.Second))
}
