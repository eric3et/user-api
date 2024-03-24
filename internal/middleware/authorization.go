package middleware

import (
	"errors"
	"net/http"
	// "net/http"
)

var ErrUnauthorized = errors.New("invalid id or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var token = r.Header.Get("Authorization")

		// if token == "" {
		// 	log.Error(ErrUnauthorized)
		// 	api.RequestErrorHandler(w, ErrUnauthorized)
		// 	return
		// }

		next.ServeHTTP(w, r)

	})
}

// Different Authorization method
// func Authorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var id string = r.URL.Query().Get("id")
// 		var token = r.Header.Get("Authorization")

// 		if id == "" || token == "" {
// 			log.Error(ErrUnauthorized)
// 			api.RequestErrorHandler(w, ErrUnauthorized)
// 			return
// 		}

// 		id_int, _ := strconv.Atoi(id)
// 		var user *tools.User = tools.GetItem(id_int)

// 		if user == nil || (token != (*user).Token) {
// 			log.Error(ErrUnauthorized)
// 			api.RequestErrorHandler(w, ErrUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)

// 	})
// }
