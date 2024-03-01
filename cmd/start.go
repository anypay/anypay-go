// Package cmd includes all the CLI commands.
package cmd

import (
	"fmt"

	"github.com/anypay/anypay-go/server"
	"github.com/spf13/cobra"
)

var cfgFile string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the web server",
	Long:  `Starts the web server with Gin and serves metrics with Prometheus.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server...")
		server.StartServer(cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here we define the command line flags for the start command
	startCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")
}
