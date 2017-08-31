package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("../service_account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	auth, err := app.Auth()
	if err != nil {
		panic(err)
	}
	tok, err := auth.CustomToken("some-id")
	if err != nil {
		panic(err)
	}
	log.Printf("tok = %v", tok)

	tok, err = auth.CustomTokenWithClaims("some-id", map[string]interface{}{
		"premium": true,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("tok = %v", tok)

	// ftok is user.getIdToken() response
	ftok := ``
	t, err := auth.VerifyIDToken(ftok)
	if err != nil {
		panic(err)
	}
	log.Printf("t = %#v", t)
}
