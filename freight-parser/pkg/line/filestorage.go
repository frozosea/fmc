package line

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"io"
	"os"
	"time"
)

const bucketName = "linesimagebucket"

type IFileStorage interface {
	UploadImage(ctx context.Context, image io.Reader) (string, error)
}

type fileStorage struct {
	client *s3.Client
}

//TODO add to .env file
func (s *fileStorage) getObjectUrl(objectKey string) string {
	s3Url := os.Getenv("S3_URL")
	return fmt.Sprintf(`%s/%s/%s`, s3Url, bucketName, objectKey)
}
func (s *fileStorage) UploadImage(ctx context.Context, image io.Reader) (string, error) {
	uniqueId := uuid.New().String()
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(uniqueId),
		ACL:    "",
		Body:   image,
	})
	return s.getObjectUrl(uniqueId), err
}

func NewFileStorage(cli *s3.Client) *fileStorage {
	return &fileStorage{client: cli}
}

type credentials struct {
}

func (s *credentials) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     os.Getenv("S3_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("S3_SECRET_ACCESS_KEY"),
		SessionToken:    "",
		Source:          "",
		CanExpire:       false,
		Expires:         time.Time{},
	}, nil
}
func ConfigureFileStorage() *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})
	_, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	cfg := aws.Config{
		Region:                      "ru-central1",
		Credentials:                 &credentials{},
		HTTPClient:                  nil,
		EndpointResolver:            nil,
		EndpointResolverWithOptions: customResolver,
		RetryMaxAttempts:            3,
		RetryMode:                   "",
		Retryer:                     nil,
		ConfigSources:               nil,
		APIOptions:                  nil,
		Logger:                      nil,
		ClientLogMode:               0,
		DefaultsMode:                "",
		RuntimeEnvironment:          aws.RuntimeEnvironment{},
	}
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(cfg)
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		panic(err)

	}
	return client
}
