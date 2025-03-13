package main

import (
	"context"

	"github.com/ruegerj/stock-sight/cmd"
	"github.com/ruegerj/stock-sight/internal/db"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			AsCommand(cmd.NewHelloCmd),
			AsCommand(cmd.NewTestDbCmd),
			fx.Annotate(
				cmd.NewRootCmd,
				fx.ParamTags(`group:"commands"`),
			),
			db.NewSQLite,
			newAppContext,
		),
		fx.NopLogger, // Disable all fx logs -> even errors
		fx.Invoke(func(*cobra.Command) {}),
	)
}

// Provides a base context for all application-scoped actions, which gets cancelled on shutdown
func newAppContext(lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})

	return ctx
}

func AsCommand(ctor any) any {
	return fx.Annotate(
		ctor,
		fx.As(new(cmd.CobraCommand)),
		fx.ResultTags(`group:"commands"`),
	)
}
