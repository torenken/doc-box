package cmd

import (
	"context"
	"errors"
	"testing"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func TestAttachCommand(t *testing.T) {

	logger := zap.S()
	ctx := context.Background()
	technicalErr := errors.New("expect Smithy http.Request value as input, got unsupported type")

	doc := domain.NewDocument("fibre-product:new-year-special", "offer")
	att := domain.NewAttachment(doc.Id, "ZGF0YQo=")
	url := "https://secret-url.de/"
	preSignedRequest := v4.PresignedHTTPRequest{URL: url, Method: "GET"}

	t.Run("error while loading document", func(t *testing.T) {
		ddb := &MockDynamo{ErrGetItem: technicalErr}
		_, err := NewAttachCommand(nil, nil, ddb, logger).Execute(ctx, att)
		assert.Error(t, err)
		assert.Equal(t, err, ErrDocLoader)
	})

	t.Run("no document found in database found", func(t *testing.T) {
		ddb := &MockDynamo{GetItemOut: GenDocItemOutput(domain.Document{})}
		_, err := NewAttachCommand(nil, nil, ddb, logger).Execute(ctx, att)
		assert.Error(t, err)
		assert.Equal(t, err, ErrDocNotFound)
	})

	t.Run("error while uploading document", func(t *testing.T) {
		ddb := &MockDynamo{GetItemOut: GenDocItemOutput(doc)}
		s3s := &MockS3{ErrPutObj: technicalErr}

		_, err := NewAttachCommand(s3s, nil, ddb, logger).Execute(ctx, att)
		assert.Error(t, err)
		assert.Equal(t, err, ErrAttUpload)
	})

	t.Run("error while creating preSignUrl", func(t *testing.T) {
		ddb := &MockDynamo{GetItemOut: GenDocItemOutput(doc)}
		s3s := &MockS3{PutObjOut: &s3.PutObjectOutput{}}
		s3p := &MockS3PreSign{ErrGetObj: technicalErr}

		_, err := NewAttachCommand(s3s, s3p, ddb, logger).Execute(ctx, att)
		assert.Error(t, err)
		assert.Equal(t, err, ErrAttPreSign)
	})

	t.Run("successfully generate the presigned url", func(t *testing.T) {
		ddb := &MockDynamo{GetItemOut: GenDocItemOutput(doc)}
		s3s := &MockS3{PutObjOut: &s3.PutObjectOutput{}}
		s3p := &MockS3PreSign{PutObjOut: &preSignedRequest}

		preSigned, err := NewAttachCommand(s3s, s3p, ddb, logger).Execute(ctx, att)
		assert.NoError(t, err)
		assert.Equal(t, preSigned.Url, url)
	})
}
