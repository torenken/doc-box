package cmd

import (
	"context"

	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func NewAttachCommand(s3s S3Putter, s3p S3PreSigner, log *zap.SugaredLogger) *AttachCommand {
	return &AttachCommand{s3s: s3s, s3p: s3p, log: log}
}

type AttachCommand struct {
	s3s S3Putter
	s3p S3PreSigner
	log *zap.SugaredLogger
}

func (c *AttachCommand) Execute(ctx context.Context, attachment *domain.Attachment) error {

	if err := NewUploader(c.s3s, c.log).Upload(ctx, *attachment); err != nil {
		c.log.Errorf("The attachment cannot be uploaded: %s", err)
		return ErrAttUpload
	}
	c.log.Infof("The upload of the attachment was successful.")

	if err := NewPreSigner(c.s3p, c.log).PreSign(ctx, attachment); err != nil {
		c.log.Errorf("The presigned url of the attachment cannot be created: %s", err)
		return ErrAttPreSign
	}

	return nil
}
