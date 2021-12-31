package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	ctx := lambdacontext.NewContext(context.TODO(), &lambdacontext.LambdaContext{
		AwsRequestID: "a3e1ca18-bca8-15ec-6519-1252bc231063",
	})

	t.Run("no request body empty", func(t *testing.T) {
		response, err := handle(ctx, events.APIGatewayProxyRequest{})
		assert.NoError(t, err)
		assert.Contains(t, response.Body, "the given body does not match to api")
	})
}
