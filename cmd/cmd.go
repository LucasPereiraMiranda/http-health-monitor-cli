package cmd

import (
	"fmt"
	"os"

	"github.com/LucasPereiraMiranda/http-health-monitor-cli/http"

	"github.com/spf13/cobra"
)

func handleCli(args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide a URL as an argument")
		return
	}

	url := args[0]
	statusCode, responseTime, isHealthy := http.CheckURL(url)

	if isHealthy {
		fmt.Printf("The URL (%s) is healthy.\nstatusCode: %d\nResponse time: %s\n", url, statusCode, responseTime)
	} else {
		fmt.Printf("The URL (%s) is not healthy.\nstatusCode: %d\nResponse time: %s\n", url, statusCode, responseTime)
	}
}

var rootCmd = &cobra.Command{
	Use:   "http-health-monitor-cli [URL]",
	Short: "Check if a URL is responding.",
	Args:  cobra.ExactArgs(1),
	Example: `
  http-health-monitor-cli https://www.google.com`,
	Run: func(cmd *cobra.Command, args []string) {
		handleCli(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
