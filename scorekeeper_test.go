package scorekeeper_test

import (
	"testing"

	"github.com/mec07/scorekeeper"
	"github.com/stretchr/testify/assert"
)

type storerMock struct {
	keyValues map[string]int
}

func newStorerMock() storerMock {
	return storerMock{
		keyValues: make(map[string]int),
	}
}

func (s storerMock) GetInt(key string) int {
	return s.keyValues[key]
}

func (s storerMock) SetInt(key string, val int) {
	s.keyValues[key] = val
}

func TestGetInitialScore(t *testing.T) {
	storer := newStorerMock()
	scorer := scorekeeper.Init(storer)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

func TestAddPoints(t *testing.T) {
	storer := newStorerMock()
	scorer := scorekeeper.Init(storer)
	assert.Equal(t, 10, scorer.Add("Sandy", 10))
	assert.Equal(t, 15, scorer.Add("Sandy", 5))
}

func TestSubtractPoints(t *testing.T) {
	storer := newStorerMock()
	scorer := scorekeeper.Init(storer)
	scorer.Add("Sandy", 10)
	assert.Equal(t, 5, scorer.Subtract("Sandy", 5))
	assert.Equal(t, 0, scorer.Subtract("Sandy", 5))
	assert.Equal(t, 0, scorer.Subtract("Sandy", 5))
}
