package shared_infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3Provider struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

func NewS3Provider(bucketName string) (*S3Provider, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error loading AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	return &S3Provider{
		client:        client,
		presignClient: presignClient,
		bucketName:    bucketName,
	}, nil
}

func (p *S3Provider) GeneratePresignedUploadURL(ctx context.Context, fileName string) (fileID string, presignedURL string, err error) {
	fileID = uuid.New().String()
	key := fmt.Sprintf("%s-%s", fileID, fileName)

	presignedRequest, err := p.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(p.bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", "", fmt.Errorf("error generating presigned URL: %w", err)
	}

	return fileID, presignedRequest.URL, nil
}
