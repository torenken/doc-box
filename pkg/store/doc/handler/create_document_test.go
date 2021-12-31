package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd"
)

func TestCreateDocumentHandler(t *testing.T) {

	ctx := context.Background()
	logger := zap.S()

	cases := []struct {
		scenario         string
		ddb              cmd.DynamoDBAPI
		req              events.APIGatewayProxyRequest
		expectStatusCode int
	}{
		{
			scenario:         "no request body empty",
			req:              events.APIGatewayProxyRequest{},
			expectStatusCode: http.StatusBadRequest,
		},
		{
			scenario:         "error on document saving",
			ddb:              &cmd.MockDynamo{ErrPut: errors.New("")},
			req:              events.APIGatewayProxyRequest{Body: "{\"name\": \"fibre-product:new-year-special\",\"type\": \"offer\"}"},
			expectStatusCode: http.StatusInternalServerError,
		},
		{
			scenario:         "document request body",
			ddb:              &cmd.MockDynamo{},
			req:              events.APIGatewayProxyRequest{Body: "{\"name\": \"fibre-product:new-year-special\",\"type\": \"offer\"}"},
			expectStatusCode: http.StatusCreated,
		},
	}

	for _, c := range cases {
		t.Run(c.scenario, func(t *testing.T) {
			resp, _ := NewCreateDocumentHandler(c.ddb, logger).Handle(ctx, c.req)
			assert.Equal(t, resp.StatusCode, c.expectStatusCode)
		})
	}
}
