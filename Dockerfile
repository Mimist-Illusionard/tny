FROM golang:1.26.1-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/url-shortener ./cmd/shortener

FROM alpine:3.22

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app
COPY --from=builder /out/url-shortener /usr/local/bin/url-shortener

USER app

EXPOSE 8082 9091

ENTRYPOINT ["url-shortener"]
CMD ["--database", "memory", "--port", "8082", "--grpc-port", "9091"]
