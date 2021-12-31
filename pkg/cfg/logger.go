package cfg

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	FieldRequestID          = "requestId"
	FieldApplicationContext = "context"
)

func NewLogger(ctx context.Context) *zap.Logger {

	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		log.Panicf("failed to get Lambda context %+v", ctx)
	}

	cfg := zap.NewProductionConfig()

	// Custom configuration
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg.InitialFields = map[string]interface{}{
		FieldRequestID:          lc.AwsRequestID,
		FieldApplicationContext: strings.ToLower(os.Getenv("APPLICATION_CONTEXT")),
	}

	// DebugEnabled is active when the DEBUG environment variable is set.
	if _, b := os.LookupEnv("DEBUG"); b {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Panic("failed to build zap logger: " + err.Error())
	}
	zap.ReplaceGlobals(logger)
	return logger
}
