package graphqlutil

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/handler"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type RegisterFieldsFunc func(query, mutation *graphql.Object)

type Config struct {
	HandlerConfig *handler.Config
}

func LoadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	config.HandlerConfig = loadEnvHandlerConfig(config.HandlerConfig)
	return config
}

func loadEnvHandlerConfig(config *handler.Config) *handler.Config {
	if config == nil {
		config = handler.NewConfig()
		web := util.GetEnvStr("GQL_WEB", "")
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

func NewHandler(config *Config, registerFields RegisterFieldsFunc) (h *handler.Handler, err error) {
	query := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{}})
	mutation := graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: graphql.Fields{}})
	registerFields(query, mutation)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: query, Mutation: mutation})
	if err != nil {
		return
	}
	config.HandlerConfig.Schema = &schema
	h = handler.New(config.HandlerConfig)
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
		return false, fmt.Errorf("MatchSelected: ResolveParams has no fields")
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
				return false, fmt.Errorf("GetSelectedFields: no fragment found with name %v", n)
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
			return false, fmt.Errorf("MatchSelected: found unexpected selection type %v", t)
		}
	}
	return false, nil
}
