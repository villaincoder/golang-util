package util

import (
	"time"
)

var LocalCN, _ = time.LoadLocation("Asia/Shanghai")
var UnixZero = time.Unix(0, 0)

func TrimTimeToDate(t time.Time) time.Time {
	ymd := t.In(LocalCN).Format("2006-01-02")
	day, _ := time.ParseInLocation("2006-01-02", ymd, LocalCN)
	return day
}

func TrimUnixMillisecondToDate(unixMillisecond int64) time.Time {
	return TrimTimeToDate(UnixMillisecondToTime(unixMillisecond))
}

func GetUnixMillisecond(time time.Time) int64 {
	millisecond := time.UnixNano() / 1e6
	if millisecond < 0 {
		return 0
	}
	return millisecond
}

func GetCurrentUnixMillisecond() int64 {
	return GetUnixMillisecond(time.Now())
}

func UnixMillisecondToTime(unixMillisecond int64) time.Time {
	if unixMillisecond <= 0 {
		return UnixZero
	}
	second := unixMillisecond / 1000
	nanoseconds := (unixMillisecond - second*1000) * 1e6
	return time.Unix(second, nanoseconds)
}

func IsInvalidTime(time time.Time) bool {
	return GetUnixMillisecond(time) == 0
}

func Ptr2Time(t *time.Time) time.Time {
	if t == nil {
		return UnixZero
	}
	return *t
}

func Time2Ptr(t time.Time) *time.Time {
	if IsInvalidTime(t) {
		return nil
	}
	return &t
}
