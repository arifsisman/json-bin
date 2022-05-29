package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	firebaseDB "firebase.google.com/go/v4/db"
	"github.com/google/uuid"

	"google.golang.org/api/option"
)

var app *firebase.App
var db *firebaseDB.Client

var err error

func InitializeApp() {
	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://jsonbin-9fae0-default-rtdb.europe-west1.firebasedatabase.app",
	}

	opt := option.WithCredentialsFile("/Users/arif/go/firebase-credentials.json")
	app, err = firebase.NewApp(ctx, config, opt)

	if err != nil {
		log.Printf("Error initializing APP: %v", err)
		return
	}

	db, err = app.Database(ctx)

	if err != nil {
		log.Printf("Error initializing DB: %v", err)
		return
	}
}

func NewBin(json string) string {
	ctx := context.Background()

	uuid := uuid.New().String()
	db.NewRef("bins").Child(uuid).Set(ctx, json)

	return uuid
}

func GetBin(uuid string) string {
	ctx := context.Background()

	var json string
	if err := db.NewRef("bins/"+uuid).Get(ctx, &json); err != nil {
		log.Fatal(err)
	}

	return json
}
