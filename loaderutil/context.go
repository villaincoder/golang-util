package loaderutil

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
)

const CtxKey = "CTX_LOADER_SERVER"

func (server *Server) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKey, &RequestHandler{
		Server:    server,
		loaderMap: make(map[string]*dataloader.Loader),
	})
}

func GetRequestHandler(ctx context.Context) (handler *RequestHandler, err error) {
	v := ctx.Value(CtxKey)
	switch v.(type) {
	case *RequestHandler:
		handler = v.(*RequestHandler)
	default:
		err = errors.New("context not found loader request handler")
	}
	return
}
