package scorekeeper_test

import (
	"testing"

	"github.com/mec07/scorekeeper"
	"github.com/stretchr/testify/assert"
)

// get the score of a player at start returns 0
func TestGetInitialScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

// add 10 to the score of a player and then get returns 10
func TestAddPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	assert.Equal(t, 10, scorer.Get("Sandy"))
}

// add 10 to the score of a player and then another 10 and get returns 20
func TestAddPointsAgain(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Add("Sandy", 10)
	assert.Equal(t, 20, scorer.Get("Sandy"))
}

func TestSubtractPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Remove("Sandy", 5)
	assert.Equal(t, 5, scorer.Get("Sandy"))
}

func TestSubtractPointsAgain(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Remove("Sandy", 5)
	scorer.Remove("Sandy", 5)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

func TestSubtractFromZero(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 0)
	scorer.Remove("Sandy", 5)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}