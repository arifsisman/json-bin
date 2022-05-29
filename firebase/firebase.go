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
var binsRef *firebaseDB.Ref

var err error

func InitializeApp() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("/Users/arif/go/firebase-credentials.json")
	config := &firebase.Config{
		DatabaseURL: "https://jsonbin-9fae0-default-rtdb.europe-west1.firebasedatabase.app",
	}
	app, err = firebase.NewApp(ctx, config, opt)

	if err != nil {
		log.Printf("error initializing app: %v", err)
		return
	}

	db, err = app.Database(ctx)

	if err != nil {
		log.Printf("error initializing DB: %v", err)
		return
	}

	binsRef = db.NewRef("bins")
}

func NewBin(record string) string {
	ctx := context.Background()

	uuid := uuid.New().String()
	binsRef.Child(uuid).Set(ctx, record)

	return uuid
}
