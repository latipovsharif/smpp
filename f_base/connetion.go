package f_base

import (
	"context"
	"smpp/ent"
	"smpp/rabbit"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const accountKey = "serviceAccountKey.json"

type Fbase struct {
	Ctx     context.Context
	FClient *firestore.Client
	Msg     ent.Messages
}

func FBaseCon(cacheMap *rabbit.CacheMap) error {
	fb := &Fbase{}
	fb.Ctx = context.Background()
	userKey := option.WithCredentialsFile(accountKey)
	app, err := firebase.NewApp(fb.Ctx, nil, userKey)
	if err != nil {
		return err
	}
	fb.Msg = ent.Messages{}
	client, err := app.Firestore(fb.Ctx)
	if err != nil {
		return err
	}
	fb.FClient = client
	go fbase(fb, cacheMap)
	return nil
}
