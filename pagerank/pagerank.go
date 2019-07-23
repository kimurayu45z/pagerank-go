package pagerank

import (
	"gonum.org/v1/gonum/mat"
)

/*
Network stands for Pagerank network
*/
type Network struct {
	n     int
	sum   map[int]float64
	links map[int]map[int]float64
}

/*
NewNetwork creates new Network
*/
func NewNetwork() *Network {
	return &Network{
		0,
		make(map[int]float64),
		make(map[int]map[int]float64),
	}
}

/*
AddLink adds link from i to j to network
*/
func (network *Network) AddLink(i int, j int, weight float64) {
	if i < 0 || j < 0 {
		panic("")
	}
	if network.n < i+1 {
		network.n = i + 1
	}
	if network.n < j+1 {
		network.n = j + 1
	}
	if network.links[i] == nil {
		network.links[i] = make(map[int]float64)
	}
	if network.links[i][j] != 0 {
		network.sum[i] -= network.links[i][j]
	}

	network.sum[i] += weight
	network.links[i][j] = weight
}

func (network *Network) initialScore() []float64 {
	value := make([]float64, network.n)
	v := float64(1) / float64(network.n)

	for i := 0; i < network.n; i++ {
		value[i] = v
	}

	return value
}

/*
Score returns the score of Pagerank.
initialScore can be nil.
exp is exponent of power method.
Pros.
- precision
Cons.
- computational complexity
- rounded error
*/
func (network *Network) Score(initialScore []float64, exp int) []float64 {
	n := network.n

	p := mat.NewDense(n, n, nil)

	for i := range network.links {
		for j := range network.links[i] {
			p.Set(i, j, network.links[i][j]/network.sum[i])
		}
	}

	if initialScore == nil {
		initialScore = network.initialScore()
	}
	s := mat.NewDense(1, n, initialScore)

	for i := 0; i < exp; i++ {
		s.Mul(s, p)
	}

	value := make([]float64, n)
	for i := 0; i < n; i++ {
		value[i] = s.At(0, i)
	}

	return value
}
