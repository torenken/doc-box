package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
	"github.com/torenken/doc-box/pkg/store/doc/handler/mapper"
)

func successfullyCreated(response interface{}, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch v := response.(type) {
	case domain.Document:
		return mapResponse(http.StatusCreated, mapper.ToDocumentResp(v, req))
	case domain.Attachment:
		return mapResponse(http.StatusCreated, mapper.ToAttachmentResp(v))
	default:
		return failed(http.StatusInternalServerError, errors.New("error during the creation response"))
	}
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
