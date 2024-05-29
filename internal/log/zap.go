package log

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ttagiyeva/superstream/internal/config"
)

var logLevels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func NewZapLogger(lc fx.Lifecycle, conf *config.Config) (*zap.Logger, error) {
	zapConf := zap.NewProductionConfig()
	zapConf.Development = conf.Logger.Env == "development"

	if level, exists := logLevels[conf.Logger.Level]; exists {
		zapConf.Level = zap.NewAtomicLevelAt(level)
	} else {
		zapConf.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger, nil
}
