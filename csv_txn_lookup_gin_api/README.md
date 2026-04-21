# csv_txn_lookup_gin_api

## Purpose

A REST API demo built with Gin for transaction lookup over HTTP.

## What this demo teaches

- Gin router and middleware patterns
- handler separation and request validation
- building a simple API for CSV-backed lookup

## Run

```bash
cd my-go-demos/csv_txn_lookup_gin_api
go run cmd/server/main.go
```

Then use the API endpoints defined in `internal/router/` to search transactions.

## Structure

- `cmd/server/` — server entrypoint
- `internal/handler/` — request handlers
- `internal/middleware/` — middleware utilities
- `internal/router/` — route definitions
- `internal/data/` — sample CSV files
- `internal/model/` — transaction model
