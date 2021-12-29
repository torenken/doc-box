package cfg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDB() *dynamodb.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return dynamodb.NewFromConfig(cfg)
}

func NewLocalDynamoDB(url string) *dynamodb.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "eu-central-1" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           url,
				SigningRegion: "eu-central-1",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	return dynamodb.NewFromConfig(cfg)
}
