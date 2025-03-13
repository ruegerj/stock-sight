package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewHelloCmd() CobraCommand {
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Says hello",
		Long:  `Dummy for checking basic functionality of cobra.`,
		Run: func(cmd *cobra.Command, args []string) {
			greeting := "hello anon"
			name := cmd.Flag("name").Value.String()
			if len(name) > 0 {
				greeting = fmt.Sprintf("hello %s", name)
			}
			fmt.Println(greeting)
		},
	}

	helloCmd.PersistentFlags().String("name", "", "Your name")

	return GenericCommand{
		cmd:  helloCmd,
		path: "root hello",
	}
}
