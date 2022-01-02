package cmd

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3PreSigner interface {
	PresignGetObject(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}
