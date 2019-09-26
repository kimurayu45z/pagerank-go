package pagerank

/*
Vector Vector
*/
type Vector = map[string]float64

/*
Matrix Matrix
*/
type Matrix = map[string]Vector

/*
GetStochastixMatrix GetStochastixMatrix
*/
func GetStochastixMatrix(linkMatrix Matrix) Matrix {
	stochasticMatrix := Matrix{}

	for src, vector := range linkMatrix {
		sum := float64(0)

		for _, value := range vector {
			sum += value
		}

		v := Vector{}

		for dst, value := range vector {
			v[dst] = value / sum
		}
		stochasticMatrix[src] = v
	}

	return stochasticMatrix
}

/*
TransitionScore TransitionScore
*/
func TransitionScore(currentScoreVector Vector, stochasticMatrix Matrix) Vector {
	score := Vector{}

	for src, vector := range stochasticMatrix {
		for dst, value := range vector {
			dstCurrentScore := score[dst]
			srcCurrentScore := currentScoreVector[src]
			score[dst] = dstCurrentScore + srcCurrentScore*value
		}
	}

	return score
}
