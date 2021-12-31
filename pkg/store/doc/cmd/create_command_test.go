package cmd

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func TestCreateCommand(t *testing.T) {

	logger := zap.S()
	ctx := context.Background()

	t.Run("error while saving document", func(t *testing.T) {
		ddb := &MockDynamo{ErrPut: errors.New("expect Smithy http.Request value as input, got unsupported type")}
		err := NewCreateCommand(ddb, logger).Execute(ctx, domain.Document{})
		assert.Error(t, err)
		assert.Equal(t, err, ErrDocCreate)
	})

	t.Run("successful saving document", func(t *testing.T) {
		ddb := &MockDynamo{PutOut: &dynamodb.PutItemOutput{}}
		err := NewCreateCommand(ddb, logger).Execute(ctx, domain.Document{})
		assert.NoError(t, err)
	})
}
