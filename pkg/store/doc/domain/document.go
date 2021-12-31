package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func NewDocument(Name, DocType string) Document {
	return Document{
		Id:             uuid.NewV4().String(),
		Name:           Name,
		Type:           DocType,
		LifecycleState: "Active",
		CreationDate:   time.Now(),
	}
}

type Document struct {
	Id             string    `dynamodbav:"id" json:"id"`
	Type           string    `dynamodbav:"type" json:"type"`
	Name           string    `dynamodbav:"name" json:"name"`
	LifecycleState string    `dynamodbav:"lifecycleState" json:"lifecycleState"`
	CreationDate   time.Time `dynamodbav:"creationDate" json:"creationDate"`
}
