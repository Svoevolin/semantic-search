package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/app"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
)

func main() {
	const op = "cmd.ui-api.main"

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.New[config.App]("../../.env")
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	container, err := app.NewContainer(ctx, cfg)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	publicServer, err := app.NewPublicServer(container).Configure(container)
	if err != nil {
		container.Logger.Panic("failed to configure public server", "error", err.Error())
	}

	application := app.New(publicServer, container)

	if err = application.Run(ctx); err != nil {
		container.Logger.Panic("failed to run app", "error", err.Error())
	}

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = application.Shutdown(ctx); err != nil {
		container.Logger.Error("failed to shutdown app", "error", err.Error())
	}
}
