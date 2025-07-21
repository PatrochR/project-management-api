package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

var ErrNoAccess = errors.New("you dont have the access")
var ErrNotFound = errors.New("not found")
var ErrDb = errors.New("database error")
var ErrWrongEmailOrPassowrd = errors.New("invalid email or password")
