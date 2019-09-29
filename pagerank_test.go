package pagerank

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStochastixMatrix(t *testing.T) {
	linkMatrix := Matrix{
		"a": {
			"b": 1,
			"c": 0.5,
			"d": 0.5,
		},
	}

	stochasticMatrix := GetStochastixMatrix(linkMatrix)
	require.Equal(t, float64(0.5), stochasticMatrix["a"]["b"])
	require.Equal(t, float64(0.25), stochasticMatrix["a"]["c"])
}

func TestTransitionScore(t *testing.T) {
	linkMatrix := Matrix{
		"a": {
			"b": 1,
			"c": 0.5,
			"d": 0.5,
		},
	}

	stochasticMatrix := GetStochastixMatrix(linkMatrix)
	score := TransitionScore(nil, stochasticMatrix)
	t.Log(score)

	for i := 0; i < 10; i++ {
		score = TransitionScore(score, stochasticMatrix)
		t.Log(score)
	}

	//require.Equal(t, 0, score)
}
