package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AllowAllCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func AllowAllHeaders(w http.ResponseWriter, r *http.Request) {
	for _, h := range r.Header["Access-Control-Request-Headers"] {
		w.Header().Set("Access-Control-Allow-Headers", h)
	}
}

var DefaultOptionsHandler = func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	AllowAllCrossOrigin(w)
	AllowAllHeaders(w, r)
}
