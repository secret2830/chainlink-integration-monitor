package daemon

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/secret2830/chainlink-integration-monitor/base"
	"github.com/secret2830/chainlink-integration-monitor/monitors/adapter"
	birita "github.com/secret2830/chainlink-integration-monitor/monitors/bsn-irita"
	"github.com/secret2830/chainlink-integration-monitor/monitors/chainlink"
	"github.com/secret2830/chainlink-integration-monitor/monitors/initiator"
)

type Application struct {
	Monitors []base.IMonitor
}

func NewApplication(config *viper.Viper) *Application {
	return &Application{
		Monitors: []base.IMonitor{
			newBIritaMonitor(config),
			newChainlinkMonitor(config),
			newInitiatorMonitor(config),
			newAdapterMonitor(config),
		},
	}
}

func (app *Application) Start() {
	for _, monitor := range app.Monitors {
		go monitor.Start()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		fmt.Printf("interrupt signal received")
	}

	logger.Info("Stopping the monitor...")

	app.Stop()
	os.Exit(0)
}

func (app *Application) Stop() {
	for _, monitor := range app.Monitors {
		monitor.Stop()
	}

	logger.Info("All monitors stopped")
}

func newBIritaMonitor(config *viper.Viper) *birita.Monitor {
	url := config.GetString("bsn-irita.endpoint")
	interval := config.GetInt64("bsn-irita.interval")
	providerAddr := config.GetString("bsn-irita.provider_addr")

	endpoint := base.NewEndpointFromURL(url)

	return birita.NewMonitor(
		endpoint,
		time.Duration(interval)*time.Second,
		providerAddr,
	)
}

func newChainlinkMonitor(config *viper.Viper) *chainlink.Monitor {
	url := config.GetString("chainlink.endpoint")
	accessKey := config.GetString("chainlink.access_key")
	secret := config.GetString("chainlink.secret")
	timeout := config.GetInt64("chainlink.timeout")
	attempts := config.GetInt("chainlink.retry")
	interval := config.GetInt64("chainlink.interval")

	endpoint := base.NewEndpoint(url, accessKey, secret)
	retry := base.NewRetryConfig(
		time.Duration(timeout)*time.Second,
		attempts,
	)

	return chainlink.NewMonitor(
		endpoint,
		retry,
		time.Duration(interval)*time.Second,
	)
}

func newInitiatorMonitor(config *viper.Viper) *initiator.Monitor {
	url := config.GetString("external-initiator.endpoint")
	timeout := config.GetInt64("external-initiator.timeout")
	attempts := config.GetInt("external-initiator.retry")
	interval := config.GetInt64("external-initiator.interval")

	endpoint := base.NewEndpointFromURL(url)
	retry := base.NewRetryConfig(
		time.Duration(timeout)*time.Second,
		attempts,
	)

	return initiator.NewMonitor(
		endpoint,
		retry,
		time.Duration(interval)*time.Second,
	)
}

func newAdapterMonitor(config *viper.Viper) *adapter.Monitor {
	url := config.GetString("external-adapter.endpoint")
	timeout := config.GetInt64("external-adapter.timeout")
	attempts := config.GetInt("external-adapter.retry")
	interval := config.GetInt64("external-adapter.interval")

	endpoint := base.NewEndpointFromURL(url)
	retry := base.NewRetryConfig(
		time.Duration(timeout)*time.Second,
		attempts,
	)

	return adapter.NewMonitor(
		endpoint,
		retry,
		time.Duration(interval)*time.Second,
	)
}
