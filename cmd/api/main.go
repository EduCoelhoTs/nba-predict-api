package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := startServer(); err != nil {
		slog.Error("error to start the server", "error", err)
		os.Exit(1)
	}
}
