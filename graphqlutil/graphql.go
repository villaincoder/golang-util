package gqlutil

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type RegisterFieldsFunc func(query, mutation *graphql.Object)

type Config struct {
	HandlerConfig *handler.Config
}

func loadEnvConfig(config *Config) *Config {
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
	config = loadEnvConfig(config)
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
