package cmd

import (
	"context"
	"fmt"

	"github.com/dennis-yeom/fw/internal/demo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command
var (
	RootCmd = &cobra.Command{
		Use:   "fw",
		Short: "Root command of the CLI",
		Long:  `This is the root command of the CLI. It serves as an entry point into other commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("For options: go run main.go -h")
		},
	}

	// SetCmd sets a key and value in Redis
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "lists contents in buckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the bucket name and endpoint from configuration
			bucket := viper.GetString("s3.bucket")
			endpoint := viper.GetString("s3.endpoint")

			if bucket == "" || endpoint == "" {
				return fmt.Errorf("bucket and endpoint must be set in the config file")
			}

			// Initialize the Demo instance with the S3 client using the bucket and endpoint
			demoInstance, err := demo.New(demo.WithS3Client(bucket))
			if err != nil {
				return fmt.Errorf("failed to initialize demo instance: %v", err)
			}

			// List files in the bucket
			if err := demoInstance.S3Client.ListFiles(context.TODO()); err != nil {
				return fmt.Errorf("failed to list files: %v", err)
			}

			return nil
		},
	}
)

// init function adds the initDemoCmd to RootCmd
func init() {
	//viper config
	viper.SetConfigName("config") // Configuration file name
	viper.SetConfigType("yaml")   // File type
	viper.AddConfigPath(".")      // Search path

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	//add commands to root command
	RootCmd.AddCommand(ListCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return RootCmd.Execute()
}
