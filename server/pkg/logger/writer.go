package logger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func getWriter(filename string, maxAgeDay int, rotationHour int) *rotatelogs.RotateLogs {
	writer, _ := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Duration(maxAgeDay)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(rotationHour)*time.Hour),
	)

	return writer
}
