# tny

![banner](documentation/banner.png)

tny is a simple URL shortener written in Go.

The service creates unique short codes for HTTP/HTTPS URLs and provides both HTTP and gRPC APIs for creating and resolving short links. It can store links either in application memory or in PostgreSQL.

## Features

- unique 10-character short codes;
- one short link per original URL;
- in-memory and PostgreSQL storage;
- HTTP server support;
- gRPC server support.

## Getting started

Clone the repository:

```bash
git clone https://github.com/Mimist-Illusionard/tny.git
cd tny
```

Install dependencies:

```bash
go mod tidy
```

Build the application:

```bash
go build -o tny ./cmd/shortener
```

After that, you can run the service with in-memory storage:

```bash
./tny --database memory --port 8082 --grpc-port 9091
```

The HTTP server will be available on port `8082` and the gRPC server on port `9091`.

## Usage

### Create a short link

Send a `POST` request to `/api/v1/links`:

```bash
curl -X POST http://localhost:8082/api/v1/links \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/page"}'
```

Example response:

```json
{
  "id": "7da40115-a3a6-4cd4-835c-e1b1c65262e1",
  "shortCode": "Abcdef123_",
  "originalUrl": "https://example.com/page"
}
```

If the same original URL is submitted again, tny returns the already existing link instead of creating another one.

### Get the original URL

Send a `GET` request to `/api/v1/links/{shortCode}`:

```bash
curl http://localhost:8082/api/v1/links/Abcdef123_
```

Example response:

```json
{
  "originalUrl": "https://example.com/page"
}
```

### Redirect to the original URL

A short code can also be opened directly:

```text
http://localhost:8082/Abcdef123_
```

tny responds with `302 Found` and redirects the client to the original URL.

## HTTP API

| Method | Endpoint                    | Description                         |
|--------|-----------------------------|-------------------------------------|
| POST   | `/api/v1/links`             | Create a short link                 |
| GET    | `/api/v1/links/{shortCode}` | Return the original URL as JSON     |
| GET    | `/{shortCode}`              | Redirect to the original URL        |

### Error responses

HTTP errors are returned as JSON:

```json
{
  "error": "invalid original url"
}
```

The API uses the following status codes:

| Status | Description                                      |
|--------|--------------------------------------------------|
| 201    | Short link created or returned                   |
| 302    | Redirect to the original URL                     |
| 400    | Invalid URL, short code or request body          |
| 404    | Short code was not found                         |
| 500    | Internal server error                            |
| 503    | A unique short code could not be generated       |

## gRPC API

tny exposes the same basic operations through gRPC on port `9091` by default.

The service definition is located in:

```text
api/v1/grpc/url.proto
```

Available methods:

```text
shortener.v1.URLShortener/CreateShortLink
shortener.v1.URLShortener/GetOriginalLink
```

Server reflection is enabled, so the API can be tested with tools such as `grpcurl`.

### Create a short link via gRPC

```bash
grpcurl -plaintext \
  -d '{"original_url":"https://example.com/page"}' \
  localhost:9091 \
  shortener.v1.URLShortener/CreateShortLink
```

### Get the original URL via gRPC

```bash
grpcurl -plaintext \
  -d '{"short_code":"Abcdef123_"}' \
  localhost:9091 \
  shortener.v1.URLShortener/GetOriginalLink
```

## Storage

tny supports two storage implementations. The implementation is selected with the `--database` flag.

### In-memory

In-memory storage requires no external dependencies:

```bash
./tny --database memory
```

All links are stored in the application process and are lost after the service stops.

### PostgreSQL

Create a `.env` file based on `.env.example`:

```env
DB_NAME=postgres
DB_HOST=localhost
DB_PORT=9090
DB_USERNAME=postgres
DB_PASSWORD=postgres
```

Then start tny with PostgreSQL storage:

```bash
./tny --database postgres --env .env
```

PostgreSQL migrations are embedded into the application and applied automatically during startup.

## Available flags

| Flag                 | Default  | Description                                |
|----------------------|----------|--------------------------------------------|
| `--database`         | `memory` | Storage implementation: `memory` or `postgres` |
| `--port`             | `8082`   | HTTP server port                           |
| `--grpc-port`        | `9091`   | gRPC server port                           |
| `--env`              | —        | Optional path to a `.env` file             |

## Docker

### Run with Docker Compose

Docker Compose starts both PostgreSQL and tny:

```bash
docker compose up --build
```

After startup:

```text
HTTP:  http://localhost:8082
gRPC:  localhost:9091
```

### Build a Docker image

```bash
docker build -t tny:local .
```

Run the image with in-memory storage:

```bash
docker run --rm \
  -p 8082:8082 \
  -p 9091:9091 \
  tny:local
```

## Taskfile

The project includes a `Taskfile.yml` for common development commands.

| Command              | Description                                      |
|----------------------|--------------------------------------------------|
| `task test`          | Run all tests                                    |
| `task build`         | Build the application                            |
| `task run:memory`    | Run HTTP and gRPC with in-memory storage         |
| `task run:postgres`  | Run HTTP and gRPC with PostgreSQL                 |
| `task proto:tools`   | Install protobuf Go generators                   |
| `task proto`         | Regenerate Go code from protobuf definitions     |
| `task docker:build`  | Build the local Docker image                     |
| `task docker:up`     | Start tny and PostgreSQL with Docker Compose      |
| `task docker:db`     | Start only PostgreSQL                            |
| `task tidy`          | Synchronize Go module dependencies               |

## Testing

Run all tests with:

```bash
go test ./...
```

or:

```bash
task test
```

The project contains tests for the service layer, HTTP and gRPC handlers, and both repository implementations.

## Project structure

```text
.
├── api/v1/grpc/                 # protobuf definition and generated gRPC code
├── cmd/shortener/               # application entry point
├── internal/
│   ├── app/                     # application startup and shutdown
│   ├── config/                  # configuration loading
│   ├── domain/                  # domain models
│   ├── repository/
│   │   ├── memory/              # in-memory repository
│   │   └── postgres/            # PostgreSQL repository and migrations
│   ├── service/                 # URL shortening business logic
│   └── transport/
│       ├── grpc/                # gRPC server and handlers
│       └── http/                # HTTP server and handlers
├── Dockerfile
├── docker-compose.yml
└── Taskfile.yml
```