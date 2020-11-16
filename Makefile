#!/usr/bin/make -f

export GO111MODULE = on
export GOPROXY = https://goproxy.io

install:
	@echo "installing chainlink-integration-monitor..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/chainlink-integration-monitor
