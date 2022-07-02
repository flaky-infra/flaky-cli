/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [id of execution]",
	Short: "The test inspect command allows you to view details relating to an execution",
	Long:  `The test inspect command allows you to view details relating to an execution.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test Inspect Called")
	},
}

func init() {
	testCmd.AddCommand(inspectCmd)
}
