package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"

	"github.com/sirupsen/logrus"
	"fasms/internal/config"
	"fasms/internal/server"
	

)

func main() {
	// Set up logrus
	logger := logrus.New()
	    // open a file
	f, err := os.OpenFile("fasms.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(f)
	logger.SetLevel(logrus.InfoLevel)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Error loading configuration")
	}

	
	// Initialize the server
	srv := server.NewServer(cfg)

	// Run the server in a goroutine
	go func() {
		logger.WithField("port", cfg.Port).Info("Starting server")
		if err := srv.Run(); err != nil {
			logger.WithError(err).Fatal("Server encountered an error")
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Allow server to clean up resources
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatal("Server forced to shut down")
	}

	logger.Info("Server stopped gracefully")
}
