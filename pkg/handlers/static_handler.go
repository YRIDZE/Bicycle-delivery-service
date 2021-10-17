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
	r.Handle("/", http.StripPrefix("/", FileServerWith404(http.Dir("./public/bicycle/dist/"), FileSystem404)))
}
