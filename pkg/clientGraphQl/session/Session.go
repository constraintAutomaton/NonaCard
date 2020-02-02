package session

import (
	"time"
)

// SessionsManager countainter of all the current session
type SessionsManager struct {
	Sessions map[int]Session
	Count    int
}

// Session session of a user containing all the token of the user, is ID and userName
type Session struct {
	Token             Token
	ThirdPartySession map[string]Token
	Manager           DBInterface
	ID                int
	UserName          string
}

// Token token of the user manage the lifecycle of the token
type Token struct {
	Value           string
	ExperiationDate time.Time
}

// DBInterface method to acces data
type DBInterface interface {
}
