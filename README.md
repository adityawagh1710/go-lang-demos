# Go Lang Demos

A small collection of Go example applications that compare CLI, concurrency, and web framework patterns.

## Project Intent

This repository is intended as a learning and portfolio showcase for Go developers. It documents several ways to build transaction lookup tooling and web services using real Go idioms.

The repo is designed to help you:

- compare CLI frameworks and argument handling (`cobra`, `flag`)
- explore concurrent CSV search implementations and worker pools
- assess web framework design with Gin and Fiber
- see a clean architecture-style web service with controller/service/repository layers

## Demos Included

### `csv_txn_lookup_cobra`
- CLI built with `cobra`
- Looks up transaction records across CSV files
- Good for learning command organization and subcommand structure

### `csv_txn_lookup_context`
- CLI built with standard `flag`
- Searches CSV files concurrently using goroutines
- Demonstrates context-aware file scanning and parallel search

### `csv_txn_lookup_flag`
- CLI built with standard `flag`
- Uses a worker pool to index and search CSV data
- Demonstrates controlling concurrency and worker assignment

### `csv_txn_lookup_gin_api`
- REST API built with Gin
- Provides transaction lookup over HTTP
- Includes middleware, routing, and handler separation

### `example4_fiber_clean_orm`
- Web service built with Fiber and GORM
- Demonstrates clean architecture layers for users, repositories, and services
- Includes request logging, error handling, and database migration

## How to Use

Each demo has its own `go.mod` and can be run independently.

Example for `csv_txn_lookup_flag`:

```bash
cd my-go-demos/csv_txn_lookup_flag
go run main.go -txn=TXN001
```

Example for `example4_fiber_clean_orm`:

```bash
cd my-go-demos/example4_fiber_clean_orm
go run main.go
```

Then open `http://localhost:3000` or follow the route definitions in `routes/`.

## Recommended Focus

To improve coherence, choose a primary narrative:

1. **Transaction lookup toolkit** – unify the CSV demos and add a shared library for parsing, searching, and benchmarking.
2. **Web service comparison** – document the Gin and Fiber APIs and build a common data model for both.
3. **Clean architecture demo** – expand `example4_fiber_clean_orm` into a complete user service with auth, database config, and tests.

## Repository Structure

- `csv_txn_lookup_cobra/` — Cobra CLI transaction lookup
- `csv_txn_lookup_context/` — concurrent CSV search using context and goroutines
- `csv_txn_lookup_flag/` — worker pool CSV search with flag-based CLI
- `csv_txn_lookup_gin_api/` — Gin-based transaction lookup API
- `example4_fiber_clean_orm/` — Fiber + GORM clean service example

## Next Steps

- Add per-demo README files with running instructions
- Add a shared `README` section that explains the intended audience and the value of each demo
- Add tests, benchmarks, and a simple comparison matrix for the key patterns shown here
