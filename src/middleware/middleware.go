package middleware

import (
	"errors"
	"fridge/src/auth"
	"net/http"

	"hodong/auth"
	"hodong/response"
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
			response.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}
