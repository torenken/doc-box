package cmd

import "go.uber.org/zap"

type DynamoDBAPI interface {
	DynamoDBPutter
}

func NewCreateCommand(ddb DynamoDBAPI, log *zap.SugaredLogger) *CreateCommand {
	return &CreateCommand{ddb: ddb, log: log}
}

type CreateCommand struct {
	ddb DynamoDBAPI
	log *zap.SugaredLogger
}
