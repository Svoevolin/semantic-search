package app

import (
	"context"
	"fmt"
)

type App struct {
	publicServer *PublicServer
	container    *Container
}

func New(publicServer *PublicServer, container *Container) *App {
	return &App{
		publicServer: publicServer,
		container:    container,
	}
}

func (app *App) Run(ctx context.Context) error {
	const op = "app.Run"

	go func() {
		if err := app.publicServer.Start(ctx); err != nil {
			app.container.Logger.Panic(op, err.Error())
		}
	}()

	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	const op = "app.Shutdown"

	if err := app.publicServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
