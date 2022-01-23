package mapper

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
	"github.com/torenken/doc-box/pkg/tmf"
)

func ToError(message string, statusCode int) tmf.Error {
	return tmf.Error{
		Code:    strconv.Itoa(statusCode),
		Message: message,
	}
}

func ToDocument(doc tmf.DocumentCreate) domain.Document {
	return domain.NewDocument(doc.GetName(), doc.GetType())
}

func ToDocumentResp(doc domain.Document, req events.APIGatewayProxyRequest) tmf.DocumentResp {
	selfLink := req.Path + "/" + doc.Id
	return tmf.DocumentResp{
		Links:          &tmf.Links{Self: tmf.NewHALSelfLink(selfLink)},
		Id:             doc.Id,
		LifecycleState: doc.LifecycleState,
		CreationDate:   &doc.CreationDate,
		Type:           &doc.Type,
		Name:           &doc.Name,
	}
}

func ToAttachmentResp(signedUrl domain.PreSignedUrl) tmf.AttachmentResp {
	return tmf.AttachmentResp{
		PreSignedUrl: &signedUrl.Url,
	}
}
