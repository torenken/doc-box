package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

var ErrAttUpload = errors.New("document attachment cannot be saved in storage")
var ErrAttDecode = errors.New("document attachment can not be decoded")

type S3Putter interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func NewUploader(s3s S3Putter, log *zap.SugaredLogger) Uploader {
	return Uploader{s3s: s3s, log: log}
}

type Uploader struct {
	s3s S3Putter
	log *zap.SugaredLogger
}

func (u Uploader) Upload(ctx context.Context, attachment domain.Attachment) error {
	u.log.Debugf("Uploading with the following input parameter: %+v", attachment)

	decodeString, err := base64.StdEncoding.DecodeString(attachment.Data)
	if err != nil {
		u.log.Errorf("attachment cannot be base64 decoded: %s", err)
		return ErrAttDecode
	}

	bucket := os.Getenv("DOCUMENT_STORAGE_NAME")
	key := fmt.Sprintf("%s.pdf", attachment.DocId) //current only one attachment per document

	u.log.Debugf("Storing attachment in %s and %s bucket.", key, bucket)

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewBuffer(decodeString),
	}

	if _, err := u.s3s.PutObject(ctx, input); err != nil {
		u.log.Errorf("attachment cannot be saved in s3 storage: %s", err)
		return ErrAttUpload
	}

	return nil
}
