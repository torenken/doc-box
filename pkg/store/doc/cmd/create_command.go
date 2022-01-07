package cmd

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

var ErrDocCreate = errors.New("technical issue while the document was created")

type DynamoDBAPI interface {
	DynamoDBPutter
}

func NewCreateCommand(ddb DynamoDBAPI, log *zap.SugaredLogger) CreateCommand {
	return CreateCommand{ddb: ddb, log: log}
}

type CreateCommand struct {
	ddb DynamoDBAPI
	log *zap.SugaredLogger
}

func (c CreateCommand) Execute(ctx context.Context, doc domain.Document) error {
	if err := NewSaver(c.ddb, c.log).Save(ctx, doc); err != nil {
		c.log.Errorf("Failed to save the document. %s", err.Error())
		return ErrDocCreate
	}
	c.log.Infof("Storing in the database was successful.")
	return nil
}
