package cfg

import (
	"context"
	"fmt"
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

var excludedEnv = map[string]bool{
	"AWS_ACCESS_KEY_ID":     true,
	"AWS_SECRET_ACCESS_KEY": true,
	"AWS_SESSION_TOKEN":     true,
}

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
		fmt.Printf("DEBUG RequestId: %s Lambda: The Log-Level was set to Debug\n", lc.AwsRequestID)
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		for _, envStr := range os.Environ() {
			fmt.Printf("DEBUG RequestId: %s Env: %s\n", lc.AwsRequestID, printEnv(envStr, excludedEnv))
		}
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Panic("failed to build zap logger: " + err.Error())
	}
	zap.ReplaceGlobals(logger)
	return logger
}

func printEnv(envStr string, excludedEnv map[string]bool) string {
	env := strings.Split(envStr, "=")
	if !excludedEnv[env[0]] {
		return envStr
	}
	return fmt.Sprintf("%s=*****<Secret>*****", env[0])
}
