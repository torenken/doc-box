package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd"
	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func NewAttachDocumentHandler(s3s cmd.S3Putter, s3p cmd.S3PreSigner, ddb cmd.DynamoDBGetter, log *zap.SugaredLogger) AttachDocumentHandler {
	return AttachDocumentHandler{s3s: s3s, s3p: s3p, ddb: ddb, log: log}
}

type AttachDocumentHandler struct {
	s3s cmd.S3Putter
	s3p cmd.S3PreSigner
	ddb cmd.DynamoDBGetter
	log *zap.SugaredLogger
}

func (h AttachDocumentHandler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.log.Info("Starting attach document")

	attachment, err := validateAndMapAttachment(req)
	if err != nil {
		h.log.Warnf("Error while mapping the request: %s", err.Error())
		return failed(http.StatusBadRequest, err)
	}
	h.log.Infof("Input validation was successful. Processing document attachment")

	preSigned, err := cmd.NewAttachCommand(h.s3s, h.s3p, h.ddb, h.log).Execute(ctx, attachment)
	if err != nil {
		if errors.Is(err, cmd.ErrDocNotFound) {
			h.log.Warnf("Failed to attach document. No document with id %s found", attachment.DocId)
			return failed(http.StatusConflict, err)
		}

		h.log.Errorf("Failed to attach document: %s", err.Error())
		return failed(http.StatusInternalServerError, err)
	}

	h.log.Infof("The attachment was created successfully with preSigned url: %s.", preSigned.Url)
	return successfullyCreated(preSigned, req)
}

func validateAndMapAttachment(req events.APIGatewayProxyRequest) (domain.Attachment, error) {
	docId := req.PathParameters["docId"]
	if len(docId) == 0 {
		return domain.Attachment{}, errors.New("the path parameter docId was missing")
	}
	return domain.NewAttachment(docId, req.Body), nil
}
