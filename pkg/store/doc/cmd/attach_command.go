package cmd

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

var ErrDocNotFound = errors.New("document cannot be found in the database")

func NewAttachCommand(s3s S3Putter, s3p S3PreSigner, ddb DynamoDBGetter, log *zap.SugaredLogger) AttachCommand {
	return AttachCommand{s3s: s3s, s3p: s3p, ddb: ddb, log: log}
}

type AttachCommand struct {
	s3s S3Putter
	s3p S3PreSigner
	ddb DynamoDBGetter
	log *zap.SugaredLogger
}

func (a AttachCommand) Execute(ctx context.Context, attachment *domain.Attachment) error {
	a.log.Infof("Perform attachment processing with documentId %s.", attachment.DocId)

	document, err := NewLoader(a.ddb, a.log).Load(ctx, attachment.DocId)
	if err != nil {
		a.log.Errorf("The document cannot be loaded: %s", err)
		return ErrDocLoader
	}

	if (domain.Document{}) == document {
		return ErrDocNotFound
	}
	a.log.Infof("The document was successful loaded.")

	if err := NewUploader(a.s3s, a.log).Upload(ctx, *attachment); err != nil {
		a.log.Errorf("The attachment cannot be uploaded: %s", err)
		return ErrAttUpload
	}
	a.log.Infof("The upload of the attachment was successful.")

	if err := NewPreSigner(a.s3p, a.log).PreSign(ctx, attachment); err != nil {
		a.log.Errorf("The presigned url of the attachment cannot be created: %s", err)
		return ErrAttPreSign
	}

	return nil
}
