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

	network := NewNetwork()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r := rand.Float64()
			if r < 0.5 {
				continue
			}
			network.AddLink(i, j, r)
		}
	}

	start := time.Now()
	rank := network.Score(nil, 10)
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds(), "seconds")

	sum := float64(0)
	for i := 0; i < n; i++ {
		sum += rank[i]
	}
	fmt.Println("sum", sum)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r := rand.Float64()
			if r < 0.9 {
				continue
			}
			network.AddLink(i, j, r)
		}
	}

	start = time.Now()
	rank = network.Score(rank, 1)
	end = time.Now()
	fmt.Println(end.Sub(start).Seconds(), "seconds")

	sum = float64(0)
	for i := 0; i < n; i++ {
		sum += rank[i]
	}
	fmt.Println("sum", sum)
}
