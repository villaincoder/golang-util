package http

import "net/http"

func AllowAllCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func AllowAllHeaders(w http.ResponseWriter, r *http.Request) {
	for _, h := range r.Header["Access-Control-Request-Headers"] {
		w.Header().Set("Access-Control-Allow-Headers", h)
	}
}
