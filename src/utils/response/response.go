package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if jsonError := json.NewEncoder(w).Encode(data); jsonError != nil {
			log.Fatal(jsonError)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, apiError error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: apiError.Error(),
	})
}
