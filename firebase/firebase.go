package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func InitializeApp() *firebase.App {
	opt := option.WithCredentialsFile("../config/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Printf("error initializing app: %v", err)
		return nil
	}

	return app
}
