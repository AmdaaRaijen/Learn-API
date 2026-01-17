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
		addr:      ":8080",
		jwtSecret: env.GetString("JWT_SECRET", ""),
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=pqgotest password=postgres dbname=pqgotest sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.db.dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	logger.Info("connected to database")

	api := api{
		config: config,
		db:     conn,
	}

	err = api.run(api.mount())

	if err != nil {
		slog.Error("Error running API", "error", err)
		os.Exit(1)
	}
}
