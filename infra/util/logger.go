package util

import (
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/config/appconfig/loglevel"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log  zerolog.Logger
	once sync.Once
)

func Logger() zerolog.Logger {
	once.Do(func() {
		var logLevel zerolog.Level

		switch config.Spec.LogLevel {
		case loglevel.Debug:
			logLevel = zerolog.DebugLevel
		case loglevel.Info:
			logLevel = zerolog.InfoLevel
		case loglevel.Warn:
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

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		Log = zerolog.New(output).
			Level(logLevel).
			With().
			Timestamp().
			Str("go_version", buildInfo.GoVersion).
			Str("git_revision", gitRevision).
			Str("version", config.Spec.Version).
			Logger()
	})

	return Log
}
