package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/watariRyo/tasktree/server/config"
	"google.golang.org/api/option"
)

type firebaseApp struct {
	*firebase.App
}

func InitFirebaseApp(ctx context.Context, cfg config.Server) (*firebaseApp, error) {
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(cfg.FirebaseSecret)))
	if err != nil {
		return nil, err
	}
	return &firebaseApp{app}, nil
}

func (app *firebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		// エラー処理
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		// エラー処理
	}
	return token, nil
}
