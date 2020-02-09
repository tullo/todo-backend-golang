package main

import (
	"net/http"
)

func cors(next http.Handler, acao string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", acao)
		w.Header().Set("access-control-allow-methods", "OPTIONS, GET, POST, PATCH, DELETE")
		w.Header().Set("access-control-allow-headers", "accept, content-type")
		if r.Method == "OPTIONS" {
			return // Preflight sets headers and we're done
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func contentTypeJSONHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func commonHandlers(next http.HandlerFunc, acao string) http.Handler {
	return contentTypeJSONHandler(cors(next, acao))
}
