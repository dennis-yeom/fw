package cmd

import (
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

	// lists all files and their versions
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "lists contents and versions in buckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the bucket name and endpoint from configuration
			bucket := viper.GetString("s3.bucket")
			endpoint := viper.GetString("s3.endpoint")

			// check if config filled
			if bucket == "" {
				return fmt.Errorf("bucket must be set in the config file")
			}
			if endpoint == "" {
				return fmt.Errorf("endpoint must be set in the config file")
			}

			// Create a new Demo instance with S3 client configuration
			d, err := demo.New(
				demo.WithS3Client(bucket, endpoint),
			)
			if err != nil {
				return fmt.Errorf("failed to configure Demo with S3 client: %v", err)
			}

			// list and err check
			if err := d.ListObjectVersions(); err != nil {
				return fmt.Errorf("failed to list object versions: %v", err)
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
