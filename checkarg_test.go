package util

import (
	"testing"
	"time"
)

type TestEnum int32

const (
	TEST1 TestEnum = iota
	TEST2
	TEST3
)

func TestCheckArgs(t *testing.T) {
	err1 := CheckArgs(map[string]Checker{
		"key": nil,
	})
	if err1 == nil {
		t.Fatal("check nil checker error")
	}
	t.Log("err1", err1)

	err2 := CheckArgs(map[string]Checker{
		"key": &StringChecker{Value: ""},
	})
	if err2 == nil {
		t.Fatal("check empty string error")
	}
	t.Log("err2", err2)

	err3 := CheckArgs(map[string]Checker{
		"key": &StringChecker{Value: "value"},
	})
	if err3 != nil {
		t.Fatal("check non-empty string error", err3)
	}

	err4 := CheckArgs(map[string]Checker{
		"key": &TimeChecker{Value: time.Time{}},
	})
	if err4 == nil {
		t.Fatal("check invalid time error")
	}
	t.Log("err4", err4)

	err5 := CheckArgs(map[string]Checker{
		"key": &TimeChecker{Value: time.Now()},
	})
	if err5 != nil {
		t.Fatal("check non-invalid time error", err5)
	}

	err6 := CheckArgs(map[string]Checker{
		"key": &UnixMillisecondChecker{Value: -10000},
	})
	if err6 == nil {
		t.Fatal("check invalid unix millisecond error")
	}
	t.Log("err6", err6)

	err7 := CheckArgs(map[string]Checker{
		"key": &UnixMillisecondChecker{Value: GetNowUnixMillisecond()},
	})
	if err7 != nil {
		t.Fatal("check now unix millisecond error", err7)
	}

	err8 := CheckArgs(map[string]Checker{
		"key": &InterfaceChecker{Value: nil},
	})
	if err8 == nil {
		t.Fatal("check nil value error")
	}
	t.Log("err8", err8)

	err9 := CheckArgs(map[string]Checker{
		"key": &InterfaceChecker{Value: &struct{}{}},
	})
	if err9 != nil {
		t.Fatal("check non-nil value error", err9)
	}

	err10 := CheckArgs(map[string]Checker{
		"key": &StringArrayChecker{Value: nil},
	})
	if err10 == nil {
		t.Fatal("check nil string array error")
	}
	t.Log("err10", err10)

	err11 := CheckArgs(map[string]Checker{
		"key": &StringArrayChecker{Value: []string{"", "1", "2"}},
	})
	if err11 == nil {
		t.Fatal("check empty item string array error")
	}
	t.Log("err11", err11)

	err12 := CheckArgs(map[string]Checker{
		"key": &StringArrayChecker{Value: []string{"1", "2", "3"}},
	})
	if err12 != nil {
		t.Fatal("check non-empty item string array error", err12)
	}

	err13 := CheckArgs(map[string]Checker{
		"key": &IntegerChecker{Value: 3, Min: 0, Max: 2},
	})
	if err13 == nil {
		t.Fatal("check too big integer error")
	}
	t.Log("err13", err13)

	err14 := CheckArgs(map[string]Checker{
		"key": &IntegerChecker{Value: -1, Min: 0, Max: 2},
	})
	if err14 == nil {
		t.Fatal("check too small integer error")
	}
	t.Log("err14", err14)

	err15 := CheckArgs(map[string]Checker{
		"key": &IntegerChecker{Value: 1, Min: 0, Max: 2},
	})
	if err15 != nil {
		t.Fatal("check integer error", err15)
	}

	err16 := CheckArgs(map[string]Checker{
		"key": &IntegerChecker{Value: int64(TestEnum(10)), Max: int64(TEST3)},
	})
	if err16 == nil {
		t.Fatal("check too big enum error")
	}
	t.Log("err16", err16)

	err17 := CheckArgs(map[string]Checker{
		"key": &IntegerChecker{Value: int64(TestEnum(-1)), Max: int64(TEST3)},
	})
	if err17 == nil {
		t.Fatal("check too small enum error")
	}
	t.Log("err17", err17)

	err18 := CheckArgs(map[string]Checker{
		"key": &IntegerChecker{Value: int64(TEST2), Max: int64(TEST3)},
	})
	if err18 != nil {
		t.Fatal("check enum error", err18)
	}

	err19 := CheckArgs(map[string]Checker{
		"key": &DecimalChecker{Value: 2.3, Min: 1.2, Max: 2.2},
	})
	if err19 == nil {
		t.Fatal("check too big decimal error")
	}
	t.Log("err19", err19)

	err20 := CheckArgs(map[string]Checker{
		"key": &DecimalChecker{Value: 1.19, Min: 1.2, Max: 2.2},
	})
	if err20 == nil {
		t.Fatal("check too small decimal error")
	}
	t.Log("err20", err20)

	err21 := CheckArgs(map[string]Checker{
		"key": &DecimalChecker{Value: 1.19, Min: 1.1, Max: 2.2},
	})
	if err21 != nil {
		t.Fatal("check decimal error", err21)
	}

	err22 := CheckArgs(map[string]Checker{
		"key": &UUIDChecker{Value: "123"},
	})
	if err22 == nil {
		t.Fatal("check invalid uuid error")
	}
	t.Log("err22", err22)

	err23 := CheckArgs(map[string]Checker{
		"key": &UUIDChecker{Value: "339dc43e-d105-472c-bb6f-e47b6d872051"},
	})
	if err23 != nil {
		t.Fatal("check uuid error", err23)
	}
}
