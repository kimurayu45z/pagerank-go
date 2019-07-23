package pagerank

import (
	"github.com/gonum/matrix/mat64"
)

func (network *Network) initialScore() []float64 {
	value := make([]float64, network.n)
	v := float64(1) / float64(network.n)

	for i := 0; i < network.n; i++ {
		value[i] = v
	}

	return value
}

/*
Score Get the score of Pagerank
initialScore can be nil
exp Exponent of power method.
The more exp the more precision for true value it has, and the more exp the more computational complexity and rouding error there is.
You can reduce computational complexity by using recently calculated result for initial score and decreasing exp.
*/
func (network *Network) Score(initialScore []float64, exp int) []float64 {
	n := network.n

	p := mat64.NewDense(n, n, nil)

	for i := range network.links {
		for j := range network.links[i] {
			p.Set(i, j, network.links[i][j]/network.sum[i])
		}
	}

	if initialScore == nil {
		initialScore = network.initialScore()
	}
	s := mat64.NewDense(1, n, initialScore)

	for i := 0; i < exp; i++ {
		s.Mul(s, p)
	}

	value := make([]float64, n)
	for i := 0; i < n; i++ {
		value[i] = s.At(0, i)
	}

	return value
}
