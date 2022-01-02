package cfg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3() *s3.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return s3.NewFromConfig(cfg)
}

func NewS3PreSign() *s3.PresignClient {
	client := NewS3()
	return s3.NewPresignClient(client)
}

func NewLocalS3(url string) *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "eu-central-1" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           url,
				SigningRegion: "eu-central-1",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	return s3.NewFromConfig(cfg)
}
