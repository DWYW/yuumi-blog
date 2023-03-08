package logger

import (
	"fmt"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	})
}

func getLevelWithName(name string) *zapcore.Level {
	alias := map[string]zapcore.Level{
		"fatal": zapcore.FatalLevel,
		"error": zapcore.ErrorLevel,
		"warn":  zapcore.WarnLevel,
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
	}
	level := alias[strings.ToLower(name)]
	return &level
}

func New(config *Config) *Logger {
	return &Logger{
		Config: *config,
	}
}

type Config struct {
	Dir          string
	MaxAgeDay    int
	RotationHour int
}

type Logger struct {
	Config
}

func (log Logger) Get(name string) *zap.Logger {
	level := getLevelWithName(name)

	if level == nil {
		return nil
	}

	filename := fmt.Sprintf("%s.log", name)
	filepath := filepath.Join(log.Dir, filename)

	return zap.New(zapcore.NewCore(
		getEncoder(),
		zapcore.AddSync(getWriter(
			filepath,
			log.MaxAgeDay,
			log.RotationHour,
		)),
		level,
	))
}
