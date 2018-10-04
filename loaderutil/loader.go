package loaderutil

import (
	"github.com/graph-gophers/dataloader"
)

type NewLoaderFunc func() *dataloader.Loader

type Server struct {
}

func NewLoaderServer() *Server {
	return &Server{}
}

func NewEmptyResults(keys dataloader.Keys) (results []*dataloader.Result, resultMap map[string]*dataloader.Result) {
	results = make([]*dataloader.Result, len(keys))
	resultMap = make(map[string]*dataloader.Result)
	for index, key := range keys {
		result := &dataloader.Result{}
		results[index] = result
		resultMap[key.String()] = result
	}
	return
}

func FillErrorResults(results *[]*dataloader.Result, len int, err error) {
	*results = make([]*dataloader.Result, len)
	for i := 0; i < len; i++ {
		(*results)[i] = &dataloader.Result{Error: err}
	}
}

type RequestHandler struct {
	Server    *Server
	loaderMap map[string]*dataloader.Loader
}

func (h *RequestHandler) GetLoader(key string, newLoader NewLoaderFunc) *dataloader.Loader {
	var loader *dataloader.Loader
	var ok bool
	if loader, ok = h.loaderMap[key]; !ok {
		loader = newLoader()
		h.loaderMap[key] = loader
	}
	return loader
}
