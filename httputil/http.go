package httputil

import (
	"net/http"

	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"

	"github.com/julienschmidt/httprouter"
)

var DefaultOptionsHandler = func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	AllowAllCrossOrigin(w)
	AllowAllHeaders(w, r)
}

type RegisterRouteFunc func(router *httprouter.Router)

type Config struct {
	Port string
}

func loadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	config.Port = util.GetEnvStr("HTTP_PORT", util.StrFallback(config.Port, "8080"))
	return config
}

func NewRouter(registerRoute RegisterRouteFunc) *httprouter.Router {
	router := httprouter.New()
	router.OPTIONS("/*path", DefaultOptionsHandler)
	registerRoute(router)
	return router
}

func Serve(config *Config, router *httprouter.Router) (err error) {
	config = loadEnvConfig(config)
	err = http.ListenAndServe(":"+config.Port, router)
	return
}

func AllowAllCrossOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func AllowAllHeaders(w http.ResponseWriter, r *http.Request) {
	for _, h := range r.Header["Access-Control-Request-Headers"] {
		w.Header().Set("Access-Control-Allow-Headers", h)
	}
}
