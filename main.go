package main

import (
	"os"

	"github.com/smartcontractkit/chainlink/core/logger"
	"go.uber.org/zap/zapcore"

	"github.com/secret2830/chainlink-integration-monitor/cmd"
)

func init() {
	logger.SetLogger(logger.CreateProductionLogger("", false, zapcore.DebugLevel, true))
}

func main() {
	rootCmd := cmd.GetRootCmd()

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
