package storage

import (
	. "../models"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
)

var (
	Client *firestore.Client
	Users  map[string]*DiscordUser
)

func loadFirestore(config []byte) error {
	if string(config) == "{}" {
		return errors.New("firebase config is empty, online storage disabled")
	}

	opt := option.WithCredentialsJSON(config)
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	Client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	Users = make(map[string]*DiscordUser)

	if err = getUsers(ctx); err != nil {
		return err
	}

	fmt.Println("Firebase loaded")
	return nil
}

func getUsers(ctx context.Context) error {
	snapshot := Client.Collection("users").Documents(ctx)
	docs, err := snapshot.GetAll()
	if err != nil {
		return err
	}
	var dUser *DiscordUser
	for _, doc := range docs {
		if err = doc.DataTo(&dUser); err != nil {
			return err
		}
		Users[doc.Ref.ID] = dUser
	}
	return nil
}
