package cfg

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {

	ctx := lambdacontext.NewContext(context.TODO(), &lambdacontext.LambdaContext{
		AwsRequestID: "a3e1ca18-bca8-15ec-6519-1252bc231063",
	})

	t.Run("no aws lambda context", func(t *testing.T) {
		assert.Panics(t, func() {
			NewLogger(context.TODO())
		})
	})

	t.Run("verify debug mode", func(t *testing.T) {
		err := os.Setenv("DEBUG", "")
		logger := NewLogger(ctx)

		assert.NoError(t, err)
		assert.True(t, logger.Core().Enabled(zap.DebugLevel))
	})

	t.Run("verify info mode", func(t *testing.T) {
		logger := NewLogger(ctx)

		assert.True(t, logger.Core().Enabled(zap.InfoLevel))
	})
}
