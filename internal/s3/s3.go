package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client encapsulates the S3 client and bucket configuration
type S3Client struct {
	client *s3.Client
	bucket string
}

// NewS3Client initializes a new S3 client for the specified bucket and endpoint
func NewS3Client(ctx context.Context, bucket, endpoint string) (*S3Client, error) {
	// Load the default configuration with a region
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	// Configure the S3 client with a custom endpoint for Linode
	s3Client := s3.New(s3.Options{
		Region:           "us-east-1",
		EndpointResolver: s3.EndpointResolverFromURL(endpoint),
		Credentials:      cfg.Credentials,
		UsePathStyle:     true, // Enable path-style addressing
	})

	return &S3Client{
		client: s3Client,
		bucket: bucket,
	}, nil
}

// ListFiles lists all files in the S3 bucket
func (s *S3Client) ListFiles(ctx context.Context) error {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
	}

	result, err := s.client.ListObjectsV2(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	fmt.Println("Files in bucket:", s.bucket)
	for _, item := range result.Contents {
		fmt.Printf(" - %s (size: %d)\n", *item.Key, item.Size)
	}
	return nil
}
