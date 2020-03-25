package utils

import (
	"encoding/json"
	"net/http"
)

// Respond responds json format with any data type
func Respond(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// RespondSuccessMessage is NOT USED at this time
func RespondSuccessMessage(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":   101,
		"result": "success",
		"msg":    message,
	}
}
