package oauthutil

import (
	"context"
	"errors"
	"net/http"
)

const CtxKey = "oauth"

func WithRequestHandler(server *Server, request *http.Request) context.Context {
	return context.WithValue(request.Context(), CtxKey, &RequestHandler{
		Server:  server,
		Request: request,
	})
}

func GetRequestHandler(ctx context.Context) (handler *RequestHandler, err error) {
	handler = ctx.Value(CtxKey).(*RequestHandler)
	if handler == nil {
		err = errors.New("context not found request handler")
	}
	return
}

func GetContextOAuthUserId(ctx context.Context) (id string, err error) {
	handler, err := GetRequestHandler(ctx)
	if err != nil {
		return
	}
	id, err = handler.GetUserId()
	return
}

func CheckContextOAuthToken(ctx context.Context) (err error) {
	handler, err := GetRequestHandler(ctx)
	if err != nil {
		return
	}
	err = handler.CheckToken()
	return
}
