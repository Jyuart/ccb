package internal

import (
	"context"
	"errors"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"golang.design/x/clipboard"
	"google.golang.org/api/option"
)

var nodePath = "ccb"

func Paste() {
	ctx := context.Background()
	dbClient, err := getDbClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	dbRef := dbClient.NewRef(nodePath)

	var clipboardContents string
	err = dbRef.Get(ctx, &clipboardContents)
	if err != nil {
		fmt.Println("There was an error getting data from the db")
		fmt.Println(err)
	}

	fmt.Println(clipboardContents)
}

func Copy() {
	clipboardContents, err := getClipboardContents()
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	dbClient, err := getDbClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	dbRef := dbClient.NewRef(nodePath)
	
	err = dbRef.Set(ctx, clipboardContents)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println("score added/updated successfully!")
}

func getClipboardContents() (string, error) {
	err := clipboard.Init()
	if err != nil {
		return "", errors.New("Error during clipboard initialization")
	}
	clipboardBytes := clipboard.Read(clipboard.FmtText)
	if clipboardBytes == nil {
		return "", errors.New("No valid text in clipboard")
	}
	return string(clipboardBytes), nil
}

func getDbClient(ctx context.Context) (*db.Client, error) {
	config := &firebase.Config{DatabaseURL: "https://coolclipboard-default-rtdb.europe-west1.firebasedatabase.app/"}
	opt := option.WithCredentialsFile("../../../../secrets/coolclipboard-firebase-adminsdk-n0dc6-da96aa0dc4.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error initializing app: %v", err))
	}
	client, err := app.Database(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error initializing db: %v", err))
	}

	return client, nil
}
