package cmd

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/cfg"
	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func TestGetterIntegration(t *testing.T) {
	if _, b := os.LookupEnv("LOCALSTACK"); !b {
		t.Skip("Start LocalStack Environment to run this test")
	}
	logger := zap.S()

	_ = os.Setenv("DOCUMENT_TABLE_NAME", "docbox-document")

	t.Run("load document by doc id", func(t *testing.T) {
		ddb := cfg.NewLocalDynamoDB("http://127.0.0.1:4566")

		doc, err := NewLoader(ddb, logger).Load(context.TODO(), "4b69ddf3-69fe-481c-a4e5-4e54456f6284")
		assert.NoError(t, err)
		assert.Equal(t, "4b69ddf3-69fe-481c-a4e5-4e54456f6284", doc.Id)
	})

	t.Run("empty document when no doc found", func(t *testing.T) {
		ddb := cfg.NewLocalDynamoDB("http://127.0.0.1:4566")

		doc, err := NewLoader(ddb, logger).Load(context.TODO(), "no-doc-id")
		assert.NoError(t, err)
		assert.Equal(t, domain.Document{}, doc)
	})
}
