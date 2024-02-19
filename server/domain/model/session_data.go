package model

type SessionData struct {
	UserID    int
	IDToken   string
	ExpiredAt int64
}
