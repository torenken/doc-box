package entity

import (
	"time"
)

type Document struct {
	Id             string    `dynamodbav:"id"`
	Type           string    `dynamodbav:"type"`
	Name           []byte    `dynamodbav:"name"`
	LifecycleState string    `dynamodbav:"lifecycleState"`
	CreationDate   time.Time `dynamodbav:"creationDate"`
}
