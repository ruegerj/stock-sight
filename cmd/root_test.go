package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestChildCommandsHaveCorrectHierarchy(t *testing.T) {
	root := &cobra.Command{Use: "acme"}
	child := &GenericCommand{
		cmd:  &cobra.Command{Use: "child"},
		path: "root child",
	}
	grandChild := &GenericCommand{
		cmd:  &cobra.Command{Use: "grandChild"},
		path: "root child grandChild",
	}

	children := []CobraCommand{child, grandChild}
	registerChildCommands(root, children)

	if child.cmd.Parent() != root {
		t.Errorf("expected child to have root as parent, got=%q", child.cmd.Parent().Use)
	}

	if grandChild.cmd.Parent() != child.cmd {
		t.Errorf("expected grand-child to have child as parent, got=%q", grandChild.cmd.Parent().Use)
	}

}
