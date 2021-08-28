package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondJsonError(w http.ResponseWriter, err *errors.RestErr) {
	RespondJson(w, err.Code, err)
}
