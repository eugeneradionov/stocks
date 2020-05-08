package logger

import "go.uber.org/zap"

type LogPreset string

const (
	Production  LogPreset = "production"
	Development LogPreset = "development"
)

func NewConfig(preset LogPreset) zap.Config {
	switch preset {
	case Production:
		return productionConfig()
	default:
		return developmentConfig()
	}
}

func developmentConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:          "json",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: false,
	}
}

func productionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
	}
}
