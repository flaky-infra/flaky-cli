/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsTestCmd = &cobra.Command{
	Use:   "ls",
	Short: "The cluster ls command allows you to view the list of last executed tests",
	Long:  `The cluster ls command allows you to view the list of last executed tests.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test Ls Called")
	},
}

func init() {
	testCmd.AddCommand(lsTestCmd)
}
