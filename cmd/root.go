/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/evandrorm89/go_stress_test/pkg/loadtester"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_stress_test",
	Short: "Execute load tests on a web service.",
	Long: `CLI tool to execute load tests on a web service.
    The user must provide the URL to test, the number of requests to perform
    and the number of multiple requests to make at a time.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		report := loadtester.RunLoadTest(url, requests, concurrency)

		cmd.Printf("Load test complete:\n")
		cmd.Printf("- Total time: %s\n", report.TotalTime)
		cmd.Printf("- Total requests: %d\n", report.TotalRequests)
		cmd.Printf("- Successful requests: %d\n", report.SuccessfulRequests)
		cmd.Printf("- Other status codes:\n")
		for code, count := range report.OtherStatusCodes {
			cmd.Printf(" - %d: %d\n", code, count)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "url to stress test")
	rootCmd.Flags().IntP("requests", "r", 1, "number of requests to perform")
	rootCmd.Flags().IntP("concurrency", "c", 1, "number of multiple requests to make at a time")
	rootCmd.MarkFlagRequired("url")
}
