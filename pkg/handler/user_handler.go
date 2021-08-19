package handler

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
	"net/http"
)

func (h Handler) Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		r := new(model.User)
		if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := h.services.Create(r)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User created"))

	default:
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
	}
}

func (h Handler) GetByEmail(w http.ResponseWriter, req *http.Request) {

}

func (h Handler) GetAll(w http.ResponseWriter, req *http.Request) {

}

func (h Handler) Update(w http.ResponseWriter, req *http.Request) {

}

func (h Handler) Delete(w http.ResponseWriter, req *http.Request) {

}
