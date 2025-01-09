package web

import (
	"encoding/json"
	"log"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err, isErr := data.(error)
	if isErr {
		data = map[string]string{"error": err.Error()}
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}
