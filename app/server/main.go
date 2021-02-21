package main

import (
	"log"
	"net/http"

	"github.com/karthikraobr/monorepo/pkg/handler"
)

func main() {
	h := handler.Handler{}
	http.Handle("/hello", h.HandleHello())
	log.Print("this is a log 2")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
