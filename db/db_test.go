package db

import (
	"testing"
)

func TestLoadEnvDBConfig(t *testing.T) {
	config := LoadEnvDBConfig(&OpenConfig{
		Debug: true,
	})
	if config.Debug != true {
		t.Fatal("config debug error")
	}
}

func TestNewPostgresDB(t *testing.T) {
	config := LoadEnvDBConfig(&OpenConfig{
		Debug: true,
	})
	db1, err1 := OpenPostgresDB(config)
	if err1 != nil {
		t.Fatal("err1", err1)
	}
	db1.Close()

	config.Port = "55432"
	_, err2 := OpenPostgresDB(config)
	if err2 == nil {
		t.Fatal("open wrong db port error")
	}
	t.Log("err2", err2)
}
