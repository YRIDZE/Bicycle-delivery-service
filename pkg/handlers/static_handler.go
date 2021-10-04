package handlers

import (
	"net/http"
)

type StaticHandler struct {
}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) RegisterRoutes(r *http.ServeMux, appH *AppHandlers) {
	r.Handle("/", http.FileServer(http.Dir("./public/bicycle/dist")))
}
