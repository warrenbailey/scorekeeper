package scorekeeper_test

import (
	"testing"

	"github.com/mec07/scorekeeper"
	"github.com/stretchr/testify/assert"
)

func TestGetInitialScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

func TestAddPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	assert.Equal(t, 10, scorer.Get("Sandy"))
}

func TestAddPointsTwice(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Add("Sandy", 5)
	assert.Equal(t, 15, scorer.Get("Sandy"))
}
