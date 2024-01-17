package main

import (
	"context"
	"log/slog"
	"os"

	api "github.com/pircuser61/go_fio/internal/api/fio_api/apiio"
	services "github.com/pircuser61/go_fio/internal/services"
	store "github.com/pircuser61/go_fio/internal/storage/postgres"
	rest "github.com/pircuser61/go_fio/internal/transport/rest"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	logger.Info("Run")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apiInstance := api.GetApi(logger)
	dbInstance := store.GetStore(logger)
	services.SetApp(apiInstance, dbInstance, logger)
	rest.RunHttpServer(ctx, logger)
}
