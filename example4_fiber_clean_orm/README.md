# example4_fiber_clean_orm

## Purpose

A Fiber and GORM demo showing a clean architecture-style web service for user management.

## What this demo teaches

- request middleware and error handling in Fiber
- repository/service/controller structure
- database connection and auto-migration with GORM

## Run

```bash
cd my-go-demos/example4_fiber_clean_orm
go run main.go
```

Then open `http://localhost:3000` and use the routes defined in `routes/`.

## Structure

- `config/` — database connection and configuration
- `controller/` — HTTP controllers
- `middleware/` — request and error middleware
- `models/` — data models
- `repository/` — data access layer
- `routes/` — route setup
- `service/` — business logic
- `utils/` — logging utility
