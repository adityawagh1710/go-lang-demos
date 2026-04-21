# csv_txn_lookup_context

## Purpose

A minimal Go CLI demo that searches transaction records in CSV files using standard `flag` parsing and concurrency.

## What this demo teaches

- `flag`-based CLI argument handling
- concurrent CSV file scanning using goroutines
- parallel search patterns with shared results

## Run

```bash
cd my-go-demos/csv_txn_lookup_context
go run main.go -txn=TXN001
```

## Structure

- `data/` — sample transaction CSV files
- `loader/` — file discovery and parallel search logic
