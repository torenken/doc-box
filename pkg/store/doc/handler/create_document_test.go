package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd"
)

func TestCreateDocumentHandler(t *testing.T) {

	ctx := context.Background()
	logger := zap.S()

	apiRequest := events.APIGatewayProxyRequest{
		Body: "{\"name\": \"fibre-product:new-year-special\",\"type\": \"offer\"}",
	}

	cases := []struct {
		scenario         string
		ddb              cmd.DynamoDBAPI
		req              events.APIGatewayProxyRequest
		expectStatusCode int
	}{
		{
			scenario:         "request body is empty",
			req:              events.APIGatewayProxyRequest{},
			expectStatusCode: http.StatusBadRequest,
		},
		{
			scenario:         "error on document saving",
			ddb:              &cmd.MockDynamo{ErrPutItem: errors.New("error from dynamodb")},
			req:              apiRequest,
			expectStatusCode: http.StatusInternalServerError,
		},
		{
			scenario:         "document request body",
			ddb:              &cmd.MockDynamo{PutItemOut: &dynamodb.PutItemOutput{}},
			req:              apiRequest,
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
