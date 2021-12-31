package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd"
	"github.com/torenken/doc-box/pkg/store/doc/domain"
	"github.com/torenken/doc-box/pkg/store/doc/handler/mapper"
	"github.com/torenken/doc-box/pkg/tmf"
)

func NewCreateDocumentHandler(ddb cmd.DynamoDBAPI, log *zap.SugaredLogger) CreateDocumentHandler {
	return CreateDocumentHandler{ddb: ddb, log: log}
}

type CreateDocumentHandler struct {
	ddb cmd.DynamoDBAPI
	log *zap.SugaredLogger
}

func (h CreateDocumentHandler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.log.Info("Starting create document")

	document, err := validateAndMap(req)
	if err != nil {
		h.log.Warnf("Error while mapping the request: %s", err.Error())
		return mapResponse(http.StatusBadRequest, mapper.ToError(err.Error(), 99))
	}
	h.log.Infof("Input validation was successful. Processing document creation")

	// Create Document
	if err := cmd.NewCreateCommand(h.ddb, h.log).Execute(ctx, document); err != nil {
		h.log.Errorf("Failed to create document: %s", err.Error())
		return mapResponse(http.StatusInternalServerError, mapper.ToError(err.Error(), 99))
	}
	return mapResponse(http.StatusCreated, mapper.ToDocumentResp(document, req))
}

func validateAndMap(req events.APIGatewayProxyRequest) (domain.Document, error) {
	var docReq tmf.DocumentCreate
	if err := json.Unmarshal([]byte(req.Body), &docReq); err != nil {
		return domain.Document{}, errors.New("the given body does not match to api. Please check the request")
	}
	return mapper.ToDocument(docReq), nil
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