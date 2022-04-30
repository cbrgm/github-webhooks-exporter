# github-webhooks-exporter

**Prometheus Exporter for Github Webhook Events**

## Installation

For pre-built binaries please take a look at the releases.  
https://github.com/cbrgm/github-webhooks-exporter/releases

### Container Usage

```bash
docker pull quay.io/cbrgm/github-webhooks-exporter:latest
docker run --rm -p 9212:9212 -p 9213:9213 quay.io/cbrgm/github-webhooks-exporter --github.webhook-secret=<id here>
```

Port `9212` is the (external) webhook receiver endpoint. Running this locally, hooks can be received at `localhost:9212/hooks`.

Port `9213` is the (internal) metrics endpoint. Running this locally, metrics can be queried at `localhost:9213/metrics`.

## Usage

```bash
Usage: github-webhooks-exporter

Flags:
  -h, --help                                 Show context-sensitive help.
      --http.webhook-addr="0.0.0.0:9212"     The address the webhook receiver is running on
      --http.exporter-addr="0.0.0.0:9213"    The address the exporter is running on
      --http.path="/metrics"                 The path metrics will be exposed at
      --github.webhook-secret=STRING         The Github webhook secret
      --log.json                             Tell the exporter to log json and not key value pairs
      --log.level="info"                     The log level to use for filtering logs
```

## Metrics

| Name                         | Type      | Cardinality   |Help
|------------------------------|-----------|---------------|----
| github_webhook_events_total  | counter   | Event  Action | Returns the total amount of webhook events by type and action (if any).
| http_request_duration_seconds  | histogram |               | HTTP Duration in seconds.
| http_requests_total  | counter   | Code  Status  | HTTP Duration in seconds.

## Development

```bash
go get -u github.com/cbrgm/github-webhooks-exporter
```

## Contributing & License

Feel free to submit changes! See
the [Contributing Guide](https://github.com/cbrgm/contributing/blob/master/CONTRIBUTING.md). This project is open-source
and is developed under the terms of
the [Apache 2.0 License](https://github.com/cbrgm/github-webhooks-exporter/blob/main/LICENSE).
