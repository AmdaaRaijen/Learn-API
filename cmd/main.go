package main

import (
"os"
"log/slog"
)

func main() {
	config := config{
	addr: ":8080",
	}

	api := api{
	config: config,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := api.run(api.mount())

	if err != nil {
		slog.Error("Error running API", "error", err)
		os.Exit(1)
	}
}
