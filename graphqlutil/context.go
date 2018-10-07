package graphqlutil

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

const CtxKey = "CTX_GRAPHQL_REQUEST_HANDLER"

func (server *Server) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKey, &RequestHandler{
		Server:    server,
		StartedAt: time.Now(),
	})
}

func GetRequestHandler(ctx context.Context) (handler *RequestHandler, err error) {
	v := ctx.Value(CtxKey)
	switch v.(type) {
	case *RequestHandler:
		handler = v.(*RequestHandler)
	default:
		err = errors.New("context not found graphql request handler")
	}
	return
}
