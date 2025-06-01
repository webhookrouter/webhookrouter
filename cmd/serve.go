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

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts http server",
	Long:  `Serve starts the HTTP server for WebhookRouter.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		// Initialize logger and configuration
		logger := logging.NewLogger("command", "prod")
		cfg := config.Load()

		// Initialize the application
		app := bootstrap.InitApplication(ctx, cfg, logger)
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
	rootCmd.AddCommand(serveCmd)
}
