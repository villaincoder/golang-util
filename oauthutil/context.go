package oauthutil

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const CtxKey = "CTX_OAUTH_REQUEST_HANDLER"

func (server *Server) NewContext(ctx context.Context, request *http.Request) context.Context {
	return context.WithValue(ctx, CtxKey, &RequestHandler{
		Server:  server,
		Request: request,
	})
}

func GetRequestHandler(ctx context.Context) (handler *RequestHandler, err error) {
	v := ctx.Value(CtxKey)
	switch v.(type) {
	case *RequestHandler:
		handler = v.(*RequestHandler)
	default:
		err = errors.New("context not found oauth request handler")
	}
	return
}

func GetContextOAuthUserId(ctx context.Context) (id string, err error) {
	handler, err := GetRequestHandler(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	id, err = handler.GetUserId()
	return
}

func CheckContextOAuthToken(ctx context.Context) (err error) {
	handler, err := GetRequestHandler(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	err = handler.CheckToken()
	return
}
