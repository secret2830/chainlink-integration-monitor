# Chainlink Integration Monitor

The daemon is intended for monitoring components related to the chainlink integration with BSN-IRITA, which include BSN-IRITA node and `Service`, the chainlink node, the external initiator and adapter.

## Get Started

### Install

```bash
make install
```

### Configuration

Configuration is required to start the chainlink integration monitor.

The default config lies in `./config/config.yaml`. The config items can be modified by demand.

### Start

```bash
chainlink-integration-monitor [config-file]
```

### Logging

Monitoring logs are stored into `log.jsonl` in the running directory.
