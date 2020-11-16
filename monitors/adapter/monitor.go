package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/secret2830/chainlink-integration-monitor/base"
	"github.com/secret2830/chainlink-integration-monitor/common"
)

var _ base.IMonitor = &Monitor{}

type Monitor struct {
	Endpoint base.Endpoint
	Retry    base.RetryConfig
	Interval time.Duration
	failed   bool
	stopped  bool
}

type HealthCheckResult struct {
	Result bool `json:"result"`
}

func NewMonitor(
	endpoint base.Endpoint,
	retry base.RetryConfig,
	interval time.Duration,
) *Monitor {
	return &Monitor{
		Endpoint: endpoint,
		Retry:    retry,
		Interval: interval,
	}
}

func (m *Monitor) Start() {
	logger.Info("External adapter monitor started")

	for {
		err := m.checkHealth()
		m.reportStatus(err)

		if !m.stopped {
			time.Sleep(m.Interval)
			continue
		}

		return
	}
}

func (m *Monitor) checkHealth() error {
	res, err := common.HttpRequestWithRetry(
		m.Endpoint.URL,
		m.Retry.Timeout,
		m.Retry.Attempts,
	)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %s", err)
	}

	var checkResult HealthCheckResult
	err = json.Unmarshal(bytes, &checkResult)
	if err != nil {
		return fmt.Errorf("failed to unmarshal result: %s", err)
	}

	if !checkResult.Result {
		return fmt.Errorf("faulty result, expected true")
	}

	return nil
}

func (m *Monitor) reportStatus(err error) {
	if err != nil && !m.failed {
		m.failed = true
		logger.Warnf("External Adapter: unable to check health status for %s, err: %s", m.Endpoint.URL, err)
	} else if err == nil && m.failed {
		m.failed = false
		logger.Infof("External Adapter: Healthy for %s", m.Endpoint.URL)
	}
}

func (m *Monitor) Stop() {
	logger.Info("External adapter monitor stopped")
	m.stopped = true
}
