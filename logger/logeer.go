package logger

import (
	"go.uber.org/zap"
)

func Load(preset LogPreset) (*zap.Logger, error) {
	loggerConfig := NewConfig(preset)

	return loggerConfig.Build()
}
