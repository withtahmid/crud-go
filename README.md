# CRUD-Go

## Copilot wrote this README.md

A minimal Go learning project demonstrating:

-   Connecting to PostgreSQL using `pgxpool`
-   Exposing a simple HTTP endpoint to list todos

## Project structure

-   `cmd/server/main.go`: Starts the HTTP server on `:8080` and registers routes
-   `internal/database/db.go`: Database connection pool (PostgreSQL)
-   `internal/handlers/todo.go`: Handler for `/list` that returns todos as JSON

## Prerequisites

-   Go 1.25+
-   PostgreSQL running locally
-   A `todos` table in your database

## Database setup

Default connection string (hardcoded in `internal/database/db.go`):

```
postgresql://postgres:postgres@localhost:5432/postgres
```

Update it if needed.

Create the table:

```
CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  complete BOOLEAN NOT NULL DEFAULT false
);
```

Insert sample data (optional):

```
INSERT INTO todos (name, complete) VALUES
  ('Learn Go', false),
  ('Build a simple API', true);
```

## Install dependencies

```
go mod tidy
```

## Run the server

```
go run ./cmd/server
```

Server logs: `Server running on :8080`

## API

-   `GET /list` → Returns all todos as JSON
    -   Response item shape: `{ "id": number, "name": string, "complete": boolean }`

### Example

```
curl -s http://localhost:8080/list | jq
```

## Notes

-   Error handling/logging is intentionally minimal for learning purposes.
-   Consider moving the DB URL to an environment variable for real projects (e.g., `DATABASE_URL`).
