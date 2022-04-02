package entity

import (
	"time"

	"github.com/google/uuid"
)

func NewDocument(Name, DocType string) Document {
	return Document{
		Id:             uuid.NewString(),
		Name:           []byte(Name),
		Type:           DocType,
		LifecycleState: "Active",
		CreationDate:   time.Now(),
	}
}

type Document struct {
	Id             string    `dynamodbav:"id"`
	Type           string    `dynamodbav:"type"`
	Name           []byte    `dynamodbav:"name"`
	LifecycleState string    `dynamodbav:"lifecycleState"`
	CreationDate   time.Time `dynamodbav:"creationDate"`
}
