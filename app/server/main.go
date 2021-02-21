package main

import (
	"net/http"

	"github.com/karthikraobr/monorepo/pkg/handler"
)

func main() {
	h := handler.Handler{}
	http.Handle("/hello", h.HandleHello())
	http.ListenAndServe(":8080", nil)
}
