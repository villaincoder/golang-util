package util

import (
	"fmt"
	"time"
)

var LocalCN, _ = time.LoadLocation("Asia/Shanghai")
var LocalGMT, _ = time.LoadLocation("GMT")
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

func GetNowUnixMillisecond() int64 {
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

func Ptr2UnixMillisecond(t *time.Time) int64 {
	return GetUnixMillisecond(Ptr2Time(t))
}

func FormatUnixMillisecond(unixMillisecond int64, layout string, local *time.Location) string {
	t := UnixMillisecondToTime(unixMillisecond)
	if local != nil {
		t = t.In(local)
	}
	return t.Format(layout)
}

func FormatMinutes(minutes uint64) string {
	hours := minutes / 60
	return fmt.Sprintf("%d:%d", hours, minutes-hours*60)
}

func DurationFallback(duration, fallback time.Duration) time.Duration {
	if duration > 0 {
		return duration
	}
	return fallback
}
