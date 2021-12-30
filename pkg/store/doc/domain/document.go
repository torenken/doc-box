package domain

import (
	uuid "github.com/satori/go.uuid"
)

func NewDocument() *Document {
	return &Document{
		Id: uuid.NewV4().String(),
	}
}

type Document struct {
	Id string `dynamodbav:"id" json:"id"`
}
