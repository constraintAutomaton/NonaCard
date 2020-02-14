package session

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SessionsManager countainter of all the current session
type SessionsManager struct {
	sessions map[string]session
	count    int
	dbEngine DBInterface
}

func (sessionsManager *SessionsManager) CreateSession() {
	token := jwt.New(jwt.SigningMethodHS512)
	SessionsManager.sessions[token] = make

}

// session session of a user containing all the token of the user, is ID and userName
type session struct {
	Token             TokenSelf
	ThirdPartySession map[string]Token
	ID                int
	UserName          string
}

// TokenSelf token of the app
type TokenSelf struct {
	Value           string
	ExperiationDate time.Time
}

// TokenAnilist token provided by anilist
type TokenAnilist struct {
	Value       string
	InitialTime time.Time
	LifeSpan    int64
}

// IsValid return true if the token is valid
func (token TokenSelf) IsValid() bool {
	return time.Now().Unix() < token.ExperiationDate.Unix()
}

// IsValid return true if the token is valid
func (token TokenAnilist) IsValid() bool {
	return token.InitialTime.Unix()+token.LifeSpan < time.Now().Unix()
}

// DBInterface method to acces data
type DBInterface interface {
}

// Token token of an API
type Token interface {
	IsValid() bool
}
