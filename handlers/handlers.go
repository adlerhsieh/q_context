package handlers

import (
	"context"
	"log"

	"google.golang.org/appengine"
)

func AnotherFunc(ctx context.Context) {
	account, err := appengine.ServiceAccount(ctx)
	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("ServiceAccount:", account)
	}
}
