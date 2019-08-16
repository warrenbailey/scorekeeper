package scorekeeper

// Scorer will keep track of the scores
type Scorer struct {
	client Storer
}

// Storer is an interface that represents talking to Redis
type Storer interface {
	GetInt(string) int
	SetInt(string, int)
}

// Init returns a scorer which talks to a redis client
func Init(client Storer) Scorer {
	return Scorer{
		client: client,
	}
}

// Get retrieves the score of the user
func (s Scorer) Get(user string) int {
	return s.client.GetInt(user)
}

// Add adds points onto a user's score
func (s Scorer) Add(user string, points int) int {
	current := s.client.GetInt(user)
	newScore := current + points
	s.client.SetInt(user, newScore)
	return newScore
}

// Subtract remove points from a user's score
func (s Scorer) Subtract(user string, points int) int {
	current := s.client.GetInt(user)
	newScore := current - points
	if newScore < 0 {
		newScore = 0
	}
	s.client.SetInt(user, newScore)
	return newScore
}
