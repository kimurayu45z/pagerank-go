package pagerank

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPagerank(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	n := 10000

	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			a[i][j] = rand.Float64()
		}
	}

	fmt.Println("n:", n)
	start := time.Now()

	rank := Score(a, 10)

	end := time.Now()

	sum := float64(0)
	for i := 0; i < n; i++ {
		sum += rank[i]
	}
	fmt.Println("sum", sum)

	fmt.Println(end.Sub(start).Seconds(), "seconds")
}
