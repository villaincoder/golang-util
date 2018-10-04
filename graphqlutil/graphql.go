package graphqlutil

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/handler"
	"github.com/pkg/errors"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git/logutil"
	"time"
)

type RegisterFieldsFunc func(query, mutation *graphql.Object)

type Config struct {
	HandlerConfig          *handler.Config
	SlowOperationThreshold time.Duration
}

type Server struct {
	Config  *Config
	Handler *handler.Handler
}

func LoadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	config.HandlerConfig = loadEnvHandlerConfig(config.HandlerConfig)
	config.SlowOperationThreshold = util.GetEnvDuration("GRAPHQL_OPERATION_THRESHOLD", util.DurationFallback(config.SlowOperationThreshold, time.Millisecond*100))
	return config
}

func loadEnvHandlerConfig(config *handler.Config) *handler.Config {
	if config == nil {
		config = handler.NewConfig()
		web := util.GetEnvStr("GRAPHQL_WEB", "")
		if web == "Playground" {
			config.Playground = true
			config.GraphiQL = false
		} else if web == "GraphiQL" {
			config.Playground = false
			config.GraphiQL = true
		} else {
			config.Playground = false
			config.GraphiQL = false
		}
	}
	return config
}

func NewServer(config *Config, registerFields RegisterFieldsFunc) (server *Server, err error) {
	if config == nil {
		config = LoadEnvConfig(config)
	}
	query := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{}})
	mutation := graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: graphql.Fields{}})
	registerFields(query, mutation)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: query, Mutation: mutation})
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	config.HandlerConfig.Schema = &schema
	config.HandlerConfig.ResultCallbackFn = resultCallback
	server = &Server{
		Config:  config,
		Handler: handler.New(config.HandlerConfig),
	}
	return
}

func MergeFields(query, mutation *graphql.Object, queryFieldMap, mutationFieldMap map[string]*graphql.Field) {
	if queryFieldMap != nil {
		for k, v := range queryFieldMap {
			query.AddFieldConfig(k, v)
		}
	}
	if mutationFieldMap != nil {
		for k, v := range mutationFieldMap {
			mutation.AddFieldConfig(k, v)
		}
	}
}

func MatchSelected(params graphql.ResolveParams, path []string) (bool, error) {
	fieldASTs := params.Info.FieldASTs
	if len(fieldASTs) == 0 {
		return false, errors.New("MatchSelected: ResolveParams has no fields")
	}
	return matchSelectedInternal(params, fieldASTs[0].SelectionSet.Selections, path, 0)
}

func matchSelectedInternal(params graphql.ResolveParams, selections []ast.Selection, path []string, pathIndex int) (bool, error) {
	for _, s := range selections {
		switch t := s.(type) {
		case *ast.Field:
			field := s.(*ast.Field)
			v := field.Name.Value
			if v == path[pathIndex] {
				if len(path) == pathIndex+1 {
					return true, nil
				} else {
					return matchSelectedInternal(params, field.SelectionSet.Selections, path, pathIndex+1)
				}
			}
			break
		case *ast.FragmentSpread:
			n := s.(*ast.FragmentSpread).Name.Value
			fragment, ok := params.Info.Fragments[n]
			if !ok {
				return false, errors.New(fmt.Sprintf("GetSelectedFields: no fragment found with name %v", n))
			}
			match, err := matchSelectedInternal(params, fragment.GetSelectionSet().Selections, path, pathIndex)
			if err != nil {
				return false, err
			}
			if match {
				return true, nil
			}
			break
		case *ast.InlineFragment:
			fragment := s.(*ast.InlineFragment)
			match, err := matchSelectedInternal(params, fragment.GetSelectionSet().Selections, path, pathIndex)
			if err != nil {
				return false, err
			}
			if match {
				return true, nil
			}
			break
		default:
			return false, errors.New(fmt.Sprintf("MatchSelected: found unexpected selection type %v", t))
		}
	}
	return false, nil
}

func resultCallback(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte) {
	h, err := GetRequestHandler(ctx)
	if err != nil {
		return
	}
	elapsedNano := h.GetElapsedNano()
	if elapsedNano > h.Server.Config.SlowOperationThreshold {
		operationName := params.OperationName
		if operationName == "" {
			operationName = "UNKNOWN"
		}
		logutil.Warnf("slow operation[%s] elapsed time %s", operationName, elapsedNano)
	}
}

type RequestHandler struct {
	Server     *Server
	StartedAt  time.Time
	FinishedAt time.Time
}

func (h *RequestHandler) GetElapsedNano() time.Duration {
	if util.IsInvalidTime(h.FinishedAt) {
		h.FinishedAt = time.Now()
	}
	return time.Duration(h.FinishedAt.UnixNano() - h.StartedAt.UnixNano())
}
