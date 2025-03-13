package cmd

import "github.com/spf13/cobra"

type CobraCommand interface {
	Command() *cobra.Command
	Path() string
}

type GenericCommand struct {
	cmd  *cobra.Command
	path string
}

func (g GenericCommand) Command() *cobra.Command {
	return g.cmd
}

func (g GenericCommand) Path() string {
	return g.path
}
