package scorekeeper

// Scorer will keep track of the scores
type Scorer struct {
	scores map[string]int
}

// NewScorer returns a scorer which talks to a redis client
func NewScorer() Scorer {
	return Scorer{
		scores: make(map[string]int),
	}
}

// Get retrieves the score of the user
func (s Scorer) Get(user string) int {
	return s.scores[user]
}

// Add adds points onto a user's score
func (s Scorer) Add(user string, points int) {
	newScore := s.scores[user] + points
	s.scores[user] = newScore
}
