package loaderutil

import (
	"context"
	"errors"
)

const CtxKey = "CTX_LOADER_SERVER"

func (server *Server) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKey, server)
}

func GetRequestHandler(ctx context.Context) (handler *RequestHandler, err error) {
	handler = ctx.Value(CtxKey).(*RequestHandler)
	if handler == nil {
		err = errors.New("context not found loader request handler")
	}
	return
}
