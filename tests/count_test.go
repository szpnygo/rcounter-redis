package tests_test

import (
	"fmt"
	"testing"
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
