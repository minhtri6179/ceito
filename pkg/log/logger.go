package log

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Logger       = zap.SugaredLogger
	LoggerConfig = zap.Config
)

func NewLogger(cfg *LoggerConfig) (*Logger, error) {

	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	core := logger.Core()
	plain := zap.New(
		core,
		zap.AddCaller(),
		zap.ErrorOutput(zapcore.AddSync(os.Stderr)),
	)

	sugar := plain.Sugar()
	return sugar, nil
}
