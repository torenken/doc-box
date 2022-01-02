package cmd

import (
	"context"

	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

type S3API interface {
	S3Putter
}

func NewAttachCommand(s3s S3API, log *zap.SugaredLogger) *AttachCommand {
	return &AttachCommand{s3s: s3s, log: log}
}

type AttachCommand struct {
	s3s S3API
	log *zap.SugaredLogger
}

func (c *AttachCommand) Execute(ctx context.Context, attachment *domain.Attachment) error {

	if err := NewUploader(c.s3s, c.log).Upload(ctx, *attachment); err != nil {
		c.log.Errorf("document attachment cannot be uploaded: %s", err)
		return ErrAttUpload
	}
	return nil
}
