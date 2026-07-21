package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Mimist-Illusionard/url-shortener/internal/app"
	"github.com/Mimist-Illusionard/url-shortener/internal/config"
)

var database string

func main() {
	flag.StringVar(&database, "database", "memory", "Database name")
	cfg := config.New(database)

	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
