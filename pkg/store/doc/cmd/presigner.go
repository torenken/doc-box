package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

var ErrAttPreSign = errors.New("preSigned url cannot be created")

type S3PreSigner interface {
	PresignGetObject(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

func NewPreSigner(s3s S3PreSigner, log *zap.SugaredLogger) PreSigner {
	return PreSigner{s3s: s3s, log: log}
}

type PreSigner struct {
	s3s S3PreSigner
	log *zap.SugaredLogger
}

func (p PreSigner) PreSign(ctx context.Context, attachment *domain.Attachment) error {
	p.log.Debugf("PreSign the attachment with the following parameter: %+v", attachment)

	input := s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("DOCUMENT_STORAGE_NAME")),
		Key:    aws.String(fmt.Sprintf("%s.pdf", attachment.DocId)), //current only one attachment per document,
	}

	resp, err := p.s3s.PresignGetObject(ctx, &input)
	if err != nil {
		p.log.Errorf("preSigned url cannot be created: %s", err)
		return ErrAttPreSign
	}
	attachment.Url = resp.URL
	p.log.Debugf("The attachment was presigned successfully: %+v", resp)

	return nil
}
