package utils

import (
	"encoding/json"
	"net/http"
)

func SendWithMessage(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}
func PreponseData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
