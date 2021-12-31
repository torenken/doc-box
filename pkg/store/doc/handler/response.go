package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
	"github.com/torenken/doc-box/pkg/store/doc/handler/mapper"
)

func successfullyCreated(doc domain.Document, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return mapResponse(http.StatusCreated, mapper.ToDocumentResp(doc, req))
}

func failed(statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	return mapResponse(statusCode, mapper.ToError(err.Error(), 99))
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
