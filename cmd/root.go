package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webhookrouter",
	Short: "The simple, privacy-first webhook router for developer teams",
	Long: `WebhookRouter is a simple, privacy-first webhook router for developer teams.
It allows you to easily manage and route webhooks from various services to your applications without compromising on privacy or security.
It is designed to be easy to use, lightweight, and efficient, making it a great choice for developers who want to focus on building their applications without worrying about webhook management.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
