package pagerank

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPagerank(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	n := 1000
	fmt.Println("n:", n)

}
