package utils

import (
	"encoding/json"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	if data != nil {
		dataJson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(dataJson)
	}
}
