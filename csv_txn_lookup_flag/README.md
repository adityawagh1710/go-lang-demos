# csv_txn_lookup_flag

## Purpose

A Go CLI demo that uses `flag` and a worker pool to index and search transaction records from CSV files.

## What this demo teaches

- `flag` argument parsing
- worker pool concurrency model
- indexing and lookups across multiple files

## Run

```bash
cd my-go-demos/csv_txn_lookup_flag
go run main.go -txn=TXN001 [-workers=10]
```

## Structure

- `data/` — sample CSV transaction files
- `loader/` — worker pool and search implementation
