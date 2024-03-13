package reporters

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func NewLogger(level string, writers ...io.Writer) *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig()),
		zapcore.NewMultiWriteSyncer(writeSyncers(writers...)...),
		zap.NewAtomicLevelAt(logLevel(level)),
	)

	return zap.New(core, zap.AddCaller())
}

func writeSyncers(writers ...io.Writer) []zapcore.WriteSyncer {
	var res []zapcore.WriteSyncer
	for _, w := range writers {
		res = append(res, zapcore.AddSync(w))
	}
	return res
}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}
}

func logLevel(level string) zapcore.Level {
	l, ok := levelMap[level]
	if !ok {
		return zapcore.InfoLevel
	}

	return l
}
