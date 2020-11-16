package common

import (
	"net/http"
	"os"
	"time"

	"github.com/sethgrid/pester"
)

func HttpRequestWithRetry(
	url string,
	timeout time.Duration,
	retries int,
) (*http.Response, error) {
	client := pester.New()

	client.Timeout = timeout
	client.MaxRetries = retries
	client.Backoff = func(retry int) time.Duration {
		return time.Duration(retry) * time.Second
	}

	return client.Get(url)
}

// MustGetHomeDir gets the user home directory
// Panic if an error occurs
func MustGetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return homeDir
}
