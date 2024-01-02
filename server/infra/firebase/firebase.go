package firebase

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/watariRyo/tasktree/server/config"
	"google.golang.org/api/option"
)

type firebaseApp struct {
	*firebase.App
}

func InitFirebaseApp(ctx context.Context, cfg config.Server) (*firebaseApp, error) {
	jsonFile, err := os.Open(cfg.FirebaseSecret)
	if err != nil {
		fmt.Println("Could not open JSON file.", err)
		return nil, err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Could not read JSON data.", err)
		return nil, err
	}

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(jsonData))
	if err != nil {
		return nil, err
	}
	return &firebaseApp{app}, nil
}

func (app *firebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
