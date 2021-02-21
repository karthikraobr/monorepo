package handler

import (
	"log"
	"net/http"
)

type Handler struct {
	log log.Logger
}

func (h *Handler) HandleHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hellko World!"))
	}
}
