# csv_txn_lookup_cobra

## Purpose

A Cobra-based CLI demo for looking up transaction records stored in CSV files.

## What this demo teaches

- `cobra` command structure and argument parsing
- CLI app organization with commands and subcommands
- CSV file scanning for transaction lookup

## Run

```bash
cd my-go-demos/csv_txn_lookup_cobra
go run main.go
```

Adjust the code or add flags in `cmd/` to support searching by transaction ID and file path.

## Structure

- `cmd/` — Cobra command implementation
- `data/` — sample CSV files
- `loader/` — CSV scanning and search logic
- `model/` — transaction data model
