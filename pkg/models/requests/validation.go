package requests

import (
	"encoding/json"
	"net/http"
)

func ValidationErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{"validationError": err.Error()})
}
