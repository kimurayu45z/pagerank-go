package pagerank

/*
Network is struct of network
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
AddLink adds link struct to network
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
