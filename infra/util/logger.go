package util

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/constants"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log  zerolog.Logger
	once sync.Once
)

type LineInfoHook struct{}

func (h LineInfoHook) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	_, file, line, ok := runtime.Caller(0)
	if ok {
		e.Str("line", fmt.Sprintf("%s:%d", file, line))
	}
}

func Logger() zerolog.Logger {
	once.Do(func() {
		var logLevel zerolog.Level

		switch config.Spec.LogLevel {
		case constants.DebugLog:
			logLevel = zerolog.DebugLevel
		case constants.InfoLog:
			logLevel = zerolog.InfoLevel
		case constants.WarnLog:
			logLevel = zerolog.WarnLevel
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if config.IsDevelopment() {
			fileLogger := &lumberjack.Logger{
				Filename:   "bda.log",
				MaxSize:    5, //
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(output, os.Stderr, fileLogger)
		}

		buildInfo, _ := debug.ReadBuildInfo()

		var lineInfoHook LineInfoHook
		Log = zerolog.New(output).
			Level(logLevel).
			Hook(lineInfoHook).
			With().
			Timestamp().
			Str("go_version", buildInfo.GoVersion).
			Str("version", config.Spec.Version).
			Logger()
	})

	return Log
}
