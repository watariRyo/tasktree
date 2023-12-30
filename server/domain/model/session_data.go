package model

import "time"

type SessionData struct {
	UserID       string
	AccessToken  string
	IdToken      string
	RefreshToken string
	ExpiredAt    time.Time
}
