package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Debug    bool
}

func LoadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	config.Host = util.GetEnvStr("DB_HOST", util.StrFallback(config.Host, "localhost"))
	config.Port = util.GetEnvStr("DB_PORT", util.StrFallback(config.Port, "5432"))
	config.User = util.GetEnvStr("DB_USER", util.StrFallback(config.User, "postgres"))
	config.Password = util.GetEnvStr("DB_PASSWORD", util.StrFallback(config.Password, "123456q"))
	config.Name = util.GetEnvStr("DB_NAME", util.StrFallback(config.Name, "postgres"))
	config.Debug = util.GetEnvStr("DB_DEBUG", "false") == "true" || config.Debug
	return config
}

func OpenPostgres(config *Config) (db *gorm.DB, err error) {
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

func ResetPostgresSchema(db *gorm.DB, schema, user string) (err error) {
	if err = db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE;", schema)).Error; err != nil {
		return
	}
	if err = db.Exec(fmt.Sprintf("CREATE SCHEMA %s;", schema)).Error; err != nil {
		return
	}
	if err = db.Exec(fmt.Sprintf("GRANT ALL ON SCHEMA %s TO %s;", schema, user)).Error; err != nil {
		return
	}
	return
}
