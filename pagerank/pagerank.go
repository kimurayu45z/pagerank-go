package pagerank

import (
	"github.com/gonum/matrix/mat64"
)

/*
Score Get the score of Pagerank
a a[i][j] is 0 if i doesn't have link to j, otherwise weighted value based on 1
exp Exponent of power method. The more exp the more precision for true value it has , and the more exp the more rouding error there is.
*/
func Score(a [][]float64, exp int) []float64 {
	n := len(a)
	rowsum := make([]float64, n)

	for i, row := range a {
		rowsum[i] = 0
		for _, column := range row {
			if column < 0 {
				panic("")
			}
			rowsum[i] += column
		}
	}

	p := mat64.NewDense(n, n, nil)

	for i := 0; i < n; i++ {
		if rowsum[i] == 0 {
			p.Set(i, i, 1)
			continue
		}
		for j := 0; j < n; j++ {
			p.Set(i, j, a[i][j]/float64(rowsum[i]))
		}
	}

	s := mat64.NewDense(1, n, nil)
	for j := 0; j < n; j++ {
		s.Set(0, j, float64(1)/float64(n))
	}
	for i := 0; i < exp; i++ {
		s.Mul(s, p)
	}

	value := make([]float64, n)
	for i := 0; i < n; i++ {
		value[i] = s.At(0, i)
	}

	return value
}
