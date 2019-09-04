package main

import (
	"log"
	"net/http"

	"github.com/adlerhsieh/q_context/handlers"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", myHandler)

	appengine.Main()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account, err := appengine.ServiceAccount(ctx)
	if err != nil {
		log.Println("[myHandler] error:", err)
	} else {
		log.Println("[myHandler] ServiceAccount:", account)
	}

	handlers.AnotherFunc(ctx)

	w.Write([]byte("ok"))
}
