package dbutil

import (
	"testing"
)

func TestOpenPostgres(t *testing.T) {
	config := loadEnvConfig(&Config{
		Debug: true,
	})
	db1, err1 := OpenPostgres(config)
	if err1 != nil {
		t.Fatal("err1", err1)
	}
	db1.Close()

	config.Port = "55432"
	_, err2 := OpenPostgres(config)
	if err2 == nil {
		t.Fatal("open wrong db port error")
	}
	t.Log("err2", err2)
}

func TestResetPostgresSchema(t *testing.T) {
	config := loadEnvConfig(&Config{
		Debug: true,
	})
	db, err := OpenPostgres(config)
	if err != nil {
		t.Fatal("open postgres error", err)
	}
	err = ResetPostgresSchema(db, "public", config.User)
	if err != nil {
		t.Fatal("reset postgres schema error", err)
	}
}
