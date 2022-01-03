package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

var ErrDocSave = errors.New("document cannot be saved in the database")
var ErrDocMarshal = errors.New("document can not be marshaled")

type DynamoDBPutter interface {
	PutItem(ctx context.Context,
		params *dynamodb.PutItemInput,
		optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func NewSaver(ddb DynamoDBPutter, log *zap.SugaredLogger) *Saver {
	return &Saver{ddb: ddb, log: log}
}

type Saver struct {
	ddb DynamoDBPutter
	log *zap.SugaredLogger
}

func (s *Saver) Save(ctx context.Context, doc domain.Document) error {
	s.log.Debugf("Save the document with the following parameter: %+v", doc)

	item, err := attributevalue.MarshalMap(doc)
	if err != nil {
		s.log.Errorf("The document can not be marshaled: %s", err.Error())
		return ErrDocMarshal
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("DOCUMENT_TABLE_NAME")),
	}

	resp, err := s.ddb.PutItem(ctx, input)
	if err != nil {
		s.log.Errorf("The document cannot be saved in the database: %s", err)
		return ErrDocSave
	}
	s.log.Debugf("The document was saved successfully: %+v", resp)

	return nil
}
