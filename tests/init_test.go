package tests_test

import (
	"fmt"
	"testing"

	"neobaran.com/neo/rcounter"
)

var rCounter *rcounter.RCounter

func init() {
	rCounter = rcounter.NewRCounter("127.0.0.1:6379", "", 0)
}

func TestInit(t *testing.T) {
	fmt.Println(rCounter.GetClient().Ping().Result())
}
