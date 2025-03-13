package main

import (
	"github.com/ruegerj/stock-sight/cmd"
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
		),
		fx.NopLogger, // Disable all fx logs -> even errors
		fx.Invoke(func(*cobra.Command) {}),
	)
}

func AsCommand(ctor any) any {
	return fx.Annotate(
		ctor,
		fx.As(new(cmd.CobraCommand)),
		fx.ResultTags(`group:"commands"`),
	)
}
