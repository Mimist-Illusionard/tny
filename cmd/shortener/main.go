package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Mimist-Illusionard/url-shortener/internal/app"
	"github.com/Mimist-Illusionard/url-shortener/internal/config"
)

var database string
var port string
var envPath string

func main() {
	flag.StringVar(&database, "database", "memory", "Database name")
	flag.StringVar(&port, "port", "8082", "Application port")
	flag.StringVar(&envPath, "env", "", "Path to env file")
	flag.Parse()

	cfg := config.New(database, port, envPath)

	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
