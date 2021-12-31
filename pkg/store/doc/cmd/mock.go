package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MockDynamo struct {
	PutOut *dynamodb.PutItemOutput
	PutInp *dynamodb.PutItemInput
	ErrPut error
}

func (m *MockDynamo) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	m.PutInp = params
	return m.PutOut, m.ErrPut
}
