package repository

import (
	"time"

	"github.com/watariRyo/tasktree/server/domain/model"
)

type RedisClient interface {
	SaveSession(sessionID string, sessionData model.SessionData, duration time.Duration) error
	GetSession(sessionID string) (*model.SessionData, error)
	DeleteSession(sessionID string) error
}
