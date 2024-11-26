# Receipt Processor

A simple receipt processor written in Go.
Exposes 2 endpoints:

- `POST /receipts/process`: Submits a receipt for processing and returns its ID
- `GET /receipts/{id}/points`: Returns the points awarded for the receipt

## Requirements

- Go 1.23.3+
- Docker

## Build

```shell
docker build -t receipt-processor .
```

## Run

```shell
docker run -p 8080:8080 receipt-processor
```

