package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/cmd/entity"
	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

var ErrDocLoader = errors.New("document cannot be retrieved from the database")
var ErrDocUnMarshal = errors.New("document can not be marshaled")

type DynamoDBGetter interface {
	GetItem(ctx context.Context,
		params *dynamodb.GetItemInput,
		optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func NewLoader(ddb DynamoDBGetter, log *zap.SugaredLogger) Loader {
	return Loader{ddb: ddb, log: log}
}

type Loader struct {
	ddb DynamoDBGetter
	log *zap.SugaredLogger
}

func (l Loader) Load(ctx context.Context, docId string) (domain.Document, error) {
	l.log.Debugf("Load the document with the following parameter: %+v", docId)

	var input = &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DOCUMENT_TABLE_NAME")),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: docId,
			},
		},
	}

	res, err := l.ddb.GetItem(ctx, input)
	if err != nil {
		l.log.Errorf("The document with id %s can not be retrieved: %s.", docId, err)
		return domain.Document{}, ErrDocLoader
	}

	var doc entity.Document
	if err := attributevalue.UnmarshalMap(res.Item, &doc); err != nil {
		l.log.Errorf("The document with id %s can not be unmarshalled: %s", docId, err)
		return domain.Document{}, ErrDocUnMarshal
	}
	document, err := entity.Decrypt(doc)
	if err != nil {
		l.log.Errorf("The document can not be decrypted: %s", err.Error())
	}

	l.log.Debugf("The document was loaded successfully: %+v", document)

	return document, nil
}
