package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/torenken/doc-box/pkg/cfg"
	"github.com/torenken/doc-box/pkg/store/doc/handler"
)

var (
	s3s *s3.Client
	s3p *s3.PresignClient
)

func init() {
	s3s = cfg.NewS3()
	s3p = cfg.NewS3PreSign()
}

func handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.NewAttachDocumentHandler(s3s, s3p, cfg.NewLogger(ctx).Sugar()).Handle(ctx, req)
}

func main() {
	lambda.Start(handle)
}
