package repository

import (
	"context"

	"firebase.google.com/go/auth"
)

type FirebaseApp interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}
