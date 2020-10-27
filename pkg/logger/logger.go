package logger

import (
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = NewLogger()

// NewLogger returns a new instance of zap sugar logger
func NewLogger() *zap.SugaredLogger {
	logger := newZap()
	defer logger.Sync()

	return logger.Sugar()
}

// Debug aliases zap.Debugw to be able to log a message
// with optional context
func Debug(msg string, args ...interface{}) {
	logger.Debugw(msg, args...)
}

// Info aliases zap.Infow to be able to log a message
// with optional context
func Info(msg string, args ...interface{}) {
	logger.Infow(msg, args...)
}

// Warn aliases zap.Warnw to be able to log a message
// with optional context
func Warn(msg string, args ...interface{}) {
	logger.Warnw(msg, args...)
}

// Error aliases zap.Errorw to be able to log a message
// with optional context
func Error(msg string, args ...interface{}) {
	logger.Errorw(msg, args...)
}

// Fatal aliases zap.Fatalw to be able to log a message
// with optional context
func Fatal(msg string, args ...interface{}) {
	logger.Fatalw(msg, args...)
}

// Debugf aliases zap.Debugf
func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

// Infof aliases zap.Infof
func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

// Warnf aliases zap.Warnf
func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

// Errorf aliases zap.Errorf
func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

// Fatalf aliases zap.Fatalf
func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args)
}

func newZap() *zap.Logger {
	// send anything above or equal to error level to stderr
	highPriority := zap.LevelEnablerFunc(func(loggingLvl zapcore.Level) bool {
		return loggingLvl >= zapcore.ErrorLevel
	})

	// send everything less than error level to stdout
	// except debug level when debugging is turned off and vice versa
	lowPriority := zap.LevelEnablerFunc(func(loggingLvl zapcore.Level) bool {
		isLessThanErr := loggingLvl < zapcore.ErrorLevel

		if isDebugMode() {
			return isLessThanErr
		}
		return isLessThanErr && loggingLvl != zapcore.DebugLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	consoleEncoder := zapcore.NewJSONEncoder(getEncoderCfg())

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	return zap.New(core,
		zap.AddCallerSkip(1),
		zap.AddCaller(),
		zap.AddStacktrace(highPriority), // add stack traces for levels above >=error only
	)
}

func getEncoderCfg() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func isDebugMode() bool {
	mode := os.Getenv("DEBUG")
	modeBool, _ := strconv.ParseBool(mode)

	return modeBool
}
