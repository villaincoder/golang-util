package oauthutil

import (
	"net/http"
	"time"

	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"

	"gopkg.in/oauth2.v3"

	"gopkg.in/go-oauth2/redis.v1"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type Config struct {
	RedisConfig     *redis.Config
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
	MockUserId      string
}

func LoadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	config.RedisConfig = loadEnvRedisConfig(config.RedisConfig)
	config.AccessTokenExp = util.GetEnvDuration("OAUTH_ACCESS_TOKEN_EXP", util.DurationFallback(config.AccessTokenExp, time.Hour*24*7))
	config.RefreshTokenExp = util.GetEnvDuration("OAUTH_REFRESH_TOKEN_EXP", util.DurationFallback(config.AccessTokenExp, time.Hour*24*30))
	config.MockUserId = util.GetEnvStr("OAUTH_MOCK_USER_ID", util.StrFallback(config.MockUserId, ""))
	return config
}

func loadEnvRedisConfig(config *redis.Config) *redis.Config {
	var defaultDB int
	if config == nil {
		config = &redis.Config{}
		defaultDB = 10
	} else {
		defaultDB = config.DB
	}
	config.Addr = util.GetEnvStr("OAUTH_REDIS_ADDR", util.StrFallback(config.Addr, "127.0.0.1:6379"))
	config.Password = util.GetEnvStr("OAUTH_REDIS_PASSWORD", util.StrFallback(config.Password, ""))
	config.DB = util.GetEnvInt("OAUTH_REDIS_DB", defaultDB)
	return config
}

type Client struct {
	Id     string
	Secret string
}

type Server struct {
	base       *server.Server
	mockUserId string
}

func (server *Server) HandleTokenRequest(w http.ResponseWriter, r *http.Request) {
	server.base.HandleTokenRequest(w, r)
}

func NewServer(config *Config, clients []Client, passwordAuthorizationHandler server.PasswordAuthorizationHandler, internalErrorHandler server.InternalErrorHandler) *Server {
	if config == nil {
		config = LoadEnvConfig(config)
	}
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(redis.NewTokenStore(config.RedisConfig))
	clientStore := store.NewClientStore()
	for _, client := range clients {
		clientStore.Set(client.Id, &models.Client{ID: client.Id, Secret: client.Secret})
	}
	manager.MapClientStorage(clientStore)
	manager.SetPasswordTokenCfg(&manage.Config{
		IsGenerateRefresh: true,
		AccessTokenExp:    config.AccessTokenExp,
		RefreshTokenExp:   config.RefreshTokenExp,
	})
	manager.SetRefreshTokenCfg(&manage.RefreshingConfig{
		IsRemoveRefreshing: true,
		IsGenerateRefresh:  true,
		AccessTokenExp:     config.AccessTokenExp,
		RefreshTokenExp:    config.RefreshTokenExp,
	})
	base := server.NewDefaultServer(manager)
	base.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.Refreshing)
	base.SetAllowGetAccessRequest(false)
	base.SetClientInfoHandler(server.ClientBasicHandler)
	base.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	base.SetInternalErrorHandler(internalErrorHandler)

	return &Server{
		base:       base,
		mockUserId: config.MockUserId,
	}
}

type RequestHandler struct {
	Server  *Server
	Request *http.Request
	userId  string
}

func (h *RequestHandler) GetUserId() (userId string, err error) {
	if h.userId != "" {
		userId = h.userId
		return
	}
	tokenInfo, err := h.Server.base.ValidationBearerToken(h.Request)
	if err != nil {
		if h.Server.mockUserId != "" {
			userId = h.Server.mockUserId
			h.userId = userId
			err = nil
		}
	} else {
		userId = tokenInfo.GetUserID()
		h.userId = userId
	}
	return
}

func (h *RequestHandler) CheckToken() (err error) {
	_, err = h.GetUserId()
	return
}
