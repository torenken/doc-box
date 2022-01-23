package handler

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd"
	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func TestAttachDocumentHandler(t *testing.T) {
	_ = os.Setenv("DOCUMENT_TABLE_NAME", "docbox-document")

	apiRequest := events.APIGatewayProxyRequest{PathParameters: map[string]string{
		"docId": "9c3af35a-0c85-4e0c-82ec-9ee5ffa337da"},
	}

	ctx := context.Background()
	logger := zap.S()

	cases := []struct {
		scenario         string
		s3s              cmd.S3Putter
		s3p              cmd.S3PreSigner
		ddb              cmd.DynamoDBGetter
		req              events.APIGatewayProxyRequest
		expectStatusCode int
	}{
		{
			scenario:         "no path parameter docId exists",
			req:              events.APIGatewayProxyRequest{},
			expectStatusCode: http.StatusBadRequest,
		},
		{
			scenario:         "no document found in database",
			req:              apiRequest,
			ddb:              &cmd.MockDynamo{GetItemOut: cmd.GenDocItemOutput(domain.Document{})},
			expectStatusCode: http.StatusConflict,
		},
		{
			scenario:         "error while loading document from database",
			req:              apiRequest,
			ddb:              &cmd.MockDynamo{ErrGetItem: errors.New("error from dynamodb")},
			expectStatusCode: http.StatusInternalServerError,
		},
		{
			scenario:         "successfully created",
			req:              apiRequest,
			ddb:              &cmd.MockDynamo{GetItemOut: cmd.GenDocItemOutput(domain.Document{Id: "9c3af35a-0c85-4e0c-82ec-9ee5ffa337da"})},
			s3s:              &cmd.MockS3{PutObjOut: &s3.PutObjectOutput{}},
			s3p:              &cmd.MockS3PreSign{PutObjOut: &v4.PresignedHTTPRequest{URL: "https://"}},
			expectStatusCode: http.StatusCreated,
		},
	}
	for _, c := range cases {
		t.Run(c.scenario, func(t *testing.T) {
			resp, _ := NewAttachDocumentHandler(c.s3s, c.s3p, c.ddb, logger).Handle(ctx, c.req)
			assert.Equal(t, resp.StatusCode, c.expectStatusCode)
		})
	}
}
