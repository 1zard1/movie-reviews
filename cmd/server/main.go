package main

import (
	"fmt"
	"movie-reviews/pkg/config"

	"go.uber.org/zap"
)

func main() {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync() // flushes buffer, if any
	logger := prdLogger.Sugar()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalw("failed to parse config", "err", err)
	}

	fmt.Printf("cfg = %+v\n", cfg)

}
