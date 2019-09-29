package pagerank

/*
Vector Vector
*/
type Vector map[string]float64

/*
Matrix Matrix
*/
type Matrix map[string]Vector

/*
Get Get
*/
func (matrix Matrix) Get(src string, dst string) float64 {
	_, ok := matrix[src]
	if !ok {
		return 0
	}

	return matrix[src][dst]
}

/*
Set Set
*/
func (matrix Matrix) Set(src string, dst string, value float64) {
	if matrix[src] == nil {
		matrix[src] = Vector{}
	}
	matrix[src][dst] = value
}

/*
GetStochastixMatrix GetStochastixMatrix
*/
func GetStochastixMatrix(linkMatrix Matrix) Matrix {
	stochasticMatrix := Matrix{}
	sum := Vector{}
	node := map[string]bool{}

	for src := range linkMatrix {
		for dst := range linkMatrix[src] {
			sum[src] += linkMatrix[src][dst]
			node[dst] = true
		}
		node[src] = true
	}

	for src := range linkMatrix {
		v := Vector{}

		for dst, value := range linkMatrix[src] {
			v[dst] = value / sum[src]

			if linkMatrix[dst] == nil && stochasticMatrix[dst] == nil {
				for n := range node {
					stochasticMatrix.Set(dst, n, float64(1)/float64(len(node)))
				}
			}
		}
		stochasticMatrix[src] = v
	}

	return stochasticMatrix
}

/*
TransitionScore TransitionScore
*/
func TransitionScore(currentScoreVector Vector, stochasticMatrix Matrix) Vector {
	if stochasticMatrix == nil || len(stochasticMatrix) == 0 {
		return Vector{}
	}
	if currentScoreVector == nil {
		currentScoreVector = Vector{}
	}
	if len(currentScoreVector) == 0 {
		s := float64(1) / float64(len(stochasticMatrix))
		for src := range stochasticMatrix {
			currentScoreVector[src] = s
		}
	}
	score := Vector{}

	for src := range stochasticMatrix {
		for dst := range stochasticMatrix[src] {
			dstCurrentScore := score[dst]
			srcCurrentScore := currentScoreVector[src]
			score[dst] = dstCurrentScore + srcCurrentScore*stochasticMatrix[src][dst]
		}
	}

	return score
}
