/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
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

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")
	helloCmd.PersistentFlags().String("name", "", "Your name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
