package util

import (
	"testing"
	"time"
)

func TestTrimTimeToDate(t *testing.T) {
	nowTime := time.Now()
	date := TrimTimeToDate(nowTime)
	if date.Format("15:04:05.999999999") != "00:00:00" {
		t.Fatal("trim time to date error")
	}
	t.Log("date", date)
}

func TestTrimUnixMillisecondToDate(t *testing.T) {
	nowTime := time.Now()
	date := TrimUnixMillisecondToDate(GetUnixMillisecond(nowTime))
	if date.Format("15:04:05.999999999") != "00:00:00" {
		t.Fatal("trim unix millisecond to date error")
	}
	t.Log("date", date)
}

func TestGetUnixMillisecond(t *testing.T) {
	ms1 := GetUnixMillisecond(time.Now())
	if ms1 <= 0 {
		t.Fatal("now time get unix millisecond error")
	}
	t.Log("ms1", ms1)
	ms2 := GetUnixMillisecond(time.Time{})
	if ms2 != 0 {
		t.Fatal("invalid time get unix millisecond error")
	}
	t.Log("ms2", ms2)
}

func TestGetNowUnixMillisecond(t *testing.T) {
	t.Log("current unix millisecond", GetNowUnixMillisecond())
}

func TestUnixMillisecondToTime(t *testing.T) {
	time1 := UnixMillisecondToTime(-10000)
	if time1 != UnixZero {
		t.Fatal("invalid unix millisecond to time error")
	}
	t.Log("time1", time1)
	nowTime := time.Now()
	time2 := UnixMillisecondToTime(GetUnixMillisecond(nowTime))
	if time2.Format("2006-01-02 15:04:05.999") != nowTime.Format("2006-01-02 15:04:05.999") {
		t.Fatal("now unix millisecond to time error")
	}
	t.Log("time2", time2)
}

func TestIsInvalidTime(t *testing.T) {
	b1 := IsInvalidTime(time.Time{})
	if b1 != true {
		t.Fatal("empty time check error")
	}
	b2 := IsInvalidTime(UnixMillisecondToTime(-10000))
	if b2 != true {
		t.Fatal("long ago time check error")
	}
	b3 := IsInvalidTime(time.Now())
	if b3 == true {
		t.Fatal("now time check error")
	}
}

func TestTime2Ptr(t *testing.T) {
	ptr1 := Time2Ptr(time.Time{})
	if ptr1 != nil {
		t.Fatal("invalid time to pointer error")
	}
	ptr2 := Time2Ptr(time.Now())
	if ptr2 == nil {
		t.Fatal("now time to pointer error")
	}
}

func TestPtr2Time(t *testing.T) {
	time1 := Ptr2Time(nil)
	if time1 != UnixZero {
		t.Fatal("nil to time error")
	}
	now := time.Now()
	time2 := Ptr2Time(&now)
	if time2 != now {
		t.Fatal("pointer to time error")
	}
}

func TestPtr2UnixMillisecond(t *testing.T) {
	nowTime := time.Now()
	if Ptr2UnixMillisecond(&nowTime) != GetUnixMillisecond(nowTime) {
		t.Fatal("Ptr2UnixMillisecond error")
	}
}

func TestFormatUnixMillisecond(t *testing.T) {
	t.Log(FormatUnixMillisecond(GetNowUnixMillisecond(), "2006-01-02 15:04:05.999999999 -0700 MST", nil))
	t.Log(FormatUnixMillisecond(GetNowUnixMillisecond(), "2006-01-02 15:04:05.999999999 -0700 MST", LocalGMT))
}

func TestFormatMinutes(t *testing.T) {
	if FormatMinutes(10) != "0:10" {
		t.Fatal("FormatMinutes 10 error")
	}
	if FormatMinutes(100) != "1:40" {
		t.Fatal("FormatMinutes 100 error")
	}
}

func TestDurationFallback(t *testing.T) {
	if DurationFallback(0, time.Hour*2) != time.Hour*2 {
		t.Fatal("DurationFallback 0 error")
	}
	if DurationFallback(time.Minute*2, time.Hour*2) != time.Minute*2 {
		t.Fatal("DurationFallback 2 minute error")
	}
}
