package cmd

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

type MockDynamo struct {
	GetItemOut *dynamodb.GetItemOutput
	GetItemIn  *dynamodb.GetItemInput
	ErrGetItem error

	PutItemOut *dynamodb.PutItemOutput
	PutItemIn  *dynamodb.PutItemInput
	ErrPutItem error
}

func (m *MockDynamo) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	m.GetItemIn = params
	return m.GetItemOut, m.ErrGetItem
}

func (m *MockDynamo) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	m.PutItemIn = params
	return m.PutItemOut, m.ErrPutItem
}

func GenDocItemOutput(doc domain.Document) *dynamodb.GetItemOutput {
	item, _ := attributevalue.MarshalMap(doc)
	return &dynamodb.GetItemOutput{
		Item: item,
	}
}

type MockS3 struct {
	PutObjOut *s3.PutObjectOutput
	PutObjIn  *s3.PutObjectInput
	ErrPutObj error
}

func (m *MockS3) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	m.PutObjIn = params
	return m.PutObjOut, m.ErrPutObj
}

type MockS3PreSign struct {
	PutObjOut *v4.PresignedHTTPRequest
	ErrGetObj error
}

func (m *MockS3PreSign) PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(options *s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return m.PutObjOut, m.ErrGetObj
}
