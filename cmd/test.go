/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/webhookrouter/webhookrouter/internal/bootstrap"
	"github.com/webhookrouter/webhookrouter/internal/config"
	"github.com/webhookrouter/webhookrouter/internal/logging"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "For internal testing purposes only",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		// Initialize logger and configuration
		logger := logging.NewLogger("command", "dev")
		cfg := config.Load()

		// Initialize the application for testing
		app := bootstrap.InitTestApplication(ctx, cfg, logger)
		logger.Info().Msg("Starting application...")
		go func() {
			if err := app.StartHTTP(); err != nil {
				log.Printf("HTTP server error: %v", err)
			}
		}()

		// Wait for shutdown signal
		<-ctx.Done()
		logger.Info().Msg("Shutting down...")
		app.Shutdown()
		time.Sleep(time.Second)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
