package main

import (
	"context"

	rest "github.com/pircuser61/go_fio/internal/transport/rest"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rest.RunHttpServer(ctx)
}
