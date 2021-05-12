package middleware

import (
	"errors"
	"net/http"

	"fridge/src/auth"
	"fridge/src/response"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.VaildToken(r)

		if nil != err {
			response.MakeJsonError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}
