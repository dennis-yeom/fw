package demo

import (
	"context"
	"fmt"

	"github.com/dennis-yeom/fw/internal/s3"
	"github.com/spf13/viper"
)

// Demo struct will hold a pointer to an S3 client
type Demo struct {
	S3Client *s3.S3Client
}

type DemoOption func(*Demo) error

// New initializes a new Demo instance with provided options
func New(opts ...DemoOption) (*Demo, error) {
	d := &Demo{}
	for _, opt := range opts {
		if err := opt(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}

// WithS3Client configures the Demo instance with an S3 client for the given bucket and endpoint
func WithS3Client(bucket string) DemoOption {
	return func(d *Demo) error {
		// Retrieve the endpoint from configuration
		endpoint := viper.GetString("s3.endpoint")
		if endpoint == "" {
			return fmt.Errorf("endpoint must be set in the config file")
		}

		// Initialize the S3 client with the bucket and endpoint
		s3Client, err := s3.NewS3Client(context.TODO(), bucket, endpoint)
		if err != nil {
			return fmt.Errorf("failed to initialize S3 client: %v", err)
		}
		d.S3Client = s3Client
		return nil
	}
}
