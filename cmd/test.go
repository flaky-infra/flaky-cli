/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "The test command allows you to view or run flakiness analysis",
	Long:  `The test command allows you to view or run flakiness analysis.`,
}

func init() {
	rootCmd.AddCommand(testCmd)
}
