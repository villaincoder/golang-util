package db

import "testing"

type TestORMStruct1 struct {
	BaseModel
	One string
}

type TestORMStruct2 struct {
	BaseModel
	Two string
}

func TestNewBaseModel(t *testing.T) {
	baseModel := NewBaseModel()
	t.Log("baseModel", baseModel)
}

func TestMigrateModels(t *testing.T) {
	config := LoadEnvDBConfig(&OpenConfig{
		Debug: true,
	})
	db, err := OpenPostgresDB(config)
	if err != nil {
		t.Fatal("open postgres db error", err)
	}

	err1 := MigrateModels(db, true, &TestORMStruct1{}, &TestORMStruct2{})
	if err1 != nil {
		t.Fatal("migrate models error", err1)
	}

	err2 := MigrateModels(db, true, nil)
	if err2 == nil {
		t.Fatal("migrate nil model error")
	}
	t.Log("err2", err2)

	db.Close()
}

func TestCreateModels(t *testing.T) {
	config := LoadEnvDBConfig(&OpenConfig{
		Debug: true,
	})
	db, err := OpenPostgresDB(config)
	if err != nil {
		t.Fatal("open postgres db error", err)
	}
	err = MigrateModels(db, true, &TestORMStruct1{}, &TestORMStruct2{})
	if err != nil {
		t.Fatal("migrate models error", err)
	}
	err = CreateModels(db, &TestORMStruct1{
		BaseModel: NewBaseModel(),
		One:       "one",
	}, &TestORMStruct2{
		BaseModel: NewBaseModel(),
		Two:       "two",
	})
	if err != nil {
		t.Fatal("create models error", err)
	}
	err = CreateModels(db, nil)
	if err == nil {
		t.Fatal("create nil model error", err)
	}
	db.Close()
}
