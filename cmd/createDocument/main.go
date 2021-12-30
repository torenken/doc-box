package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/torenken/doc-box/pkg/cfg"
	"github.com/torenken/doc-box/pkg/store/doc/handler"
)

var (
	ddb *dynamodb.Client
)

func init() {
	ddb = cfg.NewDynamoDB()
}

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.NewCreateDocumentHandler(ddb, cfg.NewLogger(ctx).Sugar()).Handle(ctx, req)
}

func main() {
	lambda.Start(handle)
}
