package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/amdaaraijen/Learn-API/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	env.LoadENV()

	config := config{
	addr: ":8080",
	db: dbConfig{
		dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=pqgotest password=postgres dbname=pqgotest sslmode=disable"),
	},
	}

	api := api{
		config: config,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	conn, DBErr := pgx.Connect(ctx, config.db.dsn)

	if DBErr != nil {
		panic(DBErr)
	}

	defer conn.Close(ctx)

	logger.Info("connected to database")

	runnerErr := api.run(api.mount())

	if runnerErr != nil {
		slog.Error("Error running API", "error", runnerErr)
		os.Exit(1)
	}
}
