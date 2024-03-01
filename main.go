package main

import (
	"github.com/anypay/anypay-go/cmd"
	"github.com/anypay/anypay-go/log"
	"go.uber.org/zap"
)

func main() {

	// Example usage
	log.Log.Info("application.started",
		// Structured context as strongly typed Field values.
		zap.String("status", "running"),
	)

	// Don't forget to flush buffered logs before your application exits.
	defer log.Log.Sync()

	cmd.Execute()
}
