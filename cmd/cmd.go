package cmd

import (
	"context"
	"fmt"

	"github.com/dennis-yeom/fw/internal/s3"
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

	// lists all files and their versions
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "lists contents and versions in buckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the bucket name and endpoint from configuration
			bucket := viper.GetString("s3.bucket")
			endpoint := viper.GetString("s3.endpoint")

			if bucket == "" || endpoint == "" {
				return fmt.Errorf("bucket and endpoint must be set in the config file")
			}

			// Initialize the S3 client with the bucket and endpoint
			s3Client, err := s3.NewS3Client(context.TODO(), bucket, endpoint)
			if err != nil {
				return fmt.Errorf("failed to initialize S3 client: %v", err)
			}

			// Retrieve and display all object versions in the bucket
			objects, err := s3Client.GetAllObjectVersions(context.TODO())
			if err != nil {
				return fmt.Errorf("failed to list object versions: %v", err)
			}

			// Display each object and its versions
			fmt.Println("Objects and their versions in bucket:", bucket)
			for _, obj := range objects {
				fmt.Printf(" - %s (version: %s)\n", obj.Key, obj.VersionID)
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
