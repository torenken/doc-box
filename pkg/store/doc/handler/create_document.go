package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd"
)

func NewCreateDocumentHandler(ddb cmd.DynamoDBAPI, log *zap.SugaredLogger) CreateDocumentHandler {
	return CreateDocumentHandler{ddb: ddb, log: log}
}

type CreateDocumentHandler struct {
	ddb cmd.DynamoDBAPI
	log *zap.SugaredLogger
}

func (h CreateDocumentHandler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.log.Infof("Starting create document")

	//TODO
	return mapResponse(http.StatusCreated, "todo")
}

func mapResponse(statusCode int, response interface{}) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(response)
	headers := map[string]string{"Content-Type": "application/json"}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}
