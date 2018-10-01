package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type OpenConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Debug    bool
}

func LoadEnvDBConfig(config *OpenConfig) *OpenConfig {
	config.Host = util.GetEnv("DB_HOST", util.StringFallback(config.Host, "localhost"))
	config.Port = util.GetEnv("DB_PORT", util.StringFallback(config.Port, "5432"))
	config.User = util.GetEnv("DB_USER", util.StringFallback(config.User, "postgres"))
	config.Password = util.GetEnv("DB_PASSWORD", util.StringFallback(config.Password, "123456q"))
	config.Name = util.GetEnv("DB_NAME", util.StringFallback(config.Name, "postgres"))
	config.Debug = util.GetEnv("DB_DEBUG", "false") == "true" || config.Debug
	return config
}

func OpenPostgresDB(config *OpenConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Name))
	if err != nil {
		if db != nil {
			db.Close()
			db = nil
		}
		return
	}
	if config.Debug {
		db = db.Debug()
	}
	return
}
