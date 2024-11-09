package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command
var (
	RootCmd = &cobra.Command{
		Use:   "batman",
		Short: "Root command of the CLI",
		Long:  `This is the root command of the CLI. It serves as an entry point into other commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cobra is up and running!")
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
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return RootCmd.Execute()
}
