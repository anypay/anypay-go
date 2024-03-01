package log

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	var err error
	Log, err = zap.NewProduction() // NewProduction returns a logger configured for production environments.

	if err != nil {
		panic(err) // In case the logger couldn't be initialized, which is critical.
	}
}
