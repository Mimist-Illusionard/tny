package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Mimist-Illusionard/url-shortener/internal/app"
	"github.com/Mimist-Illusionard/url-shortener/internal/config"
)

var database string
var httpPort string
var grpcPort string
var envPath string

func main() {
	flag.StringVar(&database, "database", "memory", "Storage implementation: memory or postgres")
	flag.StringVar(&httpPort, "port", "8082", "HTTP server port")
	flag.StringVar(&grpcPort, "grpc-port", "9091", "gRPC server port")
	flag.StringVar(&envPath, "env", "", "Optional path to .env file")
	flag.Parse()

	cfg, err := config.New(database, httpPort, grpcPort, envPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
