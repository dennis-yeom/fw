package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 interface {
	ListFiles(ctx context.Context) error
	GetObjectVersion(ctx context.Context, key string) (string, error)
	GetAllObjectVersions(ctx context.Context) ([]ObjectInfo, error)
}

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

// GetObjectVersion retrieves the metadata of an object and returns its version ID
func (s *S3Client) GetObjectVersion(ctx context.Context, key string) (string, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	result, err := s.client.HeadObject(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to get object metadata: %w", err)
	}

	versionID := aws.ToString(result.VersionId)
	fmt.Printf("File %s in bucket %s has version ID: %s\n", key, s.bucket, versionID)
	return versionID, nil
}

// ObjectInfo holds the key (filename) and version ID of an object
type ObjectInfo struct {
	Key       string
	VersionID string
}

// GetAllObjectVersions retrieves the filename and version ID for all versions of objects in the Linode S3-compatible bucket
func (s *S3Client) GetAllObjectVersions(ctx context.Context) ([]ObjectInfo, error) {
	input := &s3.ListObjectVersionsInput{
		Bucket: aws.String(s.bucket),
	}

	result, err := s.client.ListObjectVersions(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to list object versions: %w", err)
	}

	var objects []ObjectInfo
	for _, version := range result.Versions {
		objects = append(objects, ObjectInfo{
			Key:       *version.Key,
			VersionID: aws.ToString(version.VersionId), // Includes version ID directly
		})
	}

	return objects, nil
}
