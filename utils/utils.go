package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//ValueSlice slice of values
type ValueSlice []interface{}

// Logger ...
var Logger *zap.Logger

//Has checks whether value list has param "a" value.
func (l ValueSlice) Has(a interface{}) bool {
	for _, b := range l {
		if b == a {
			return true
		}
	}
	return false
}

// NewLogger initialize our golang logger
func NewLogger() {
	var (
		config zap.Config
		err    error
	)

	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "PROD":
		config = zap.NewProductionConfig()
	case "DEV":
		config = zap.NewDevelopmentConfig()
	default:
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level.SetLevel(zap.InfoLevel)
	Logger, err = config.Build()
	if err != nil {
		zap.L().Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized", zap.Error(err))
	}

	zap.ReplaceGlobals(Logger)
}

// GetEnv ...
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
