package utils

import (
	"strings"
	"time"
)

func GetCurTimeStr() string {
	currentTime := time.Now()
	formattedTime := Time2Str(&currentTime)
	return formattedTime
}

func GetCurTime() *time.Time {
	currentTime := time.Now()
	return &currentTime
}

func Time2Str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func Str2Time(str string) *time.Time {
	if str == "" {
		return nil
	}
	var layout string
	if strings.Index(str, "T") > -1 {
		layout = time.RFC3339
	} else {
		layout = "2006-01-02 15:04:05"
	}
	t, _ := time.Parse(layout, str)
	return &t
}
