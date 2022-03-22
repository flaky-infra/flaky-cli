/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	sourceUrl       string
	commitId        string
	clusterExecName string
	watchFlag       bool
	testCmd         = &cobra.Command{
		Use:   "test",
		Short: "The test command allows you to perform the flakiness analysis of a project",
		Long: `The test command allows you to carry out the flakiness analysis of a local or remote project,
correctly configured with a .flaky.yaml file, by sending it to one of the added clusters.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Test Called")
		},
	}
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&sourceUrl, "source", "s", ".", "url at which the project can be found, local path or repository url")
	testCmd.Flags().StringVarP(&commitId, "commitId", "i", "", "hash of the commit to refer to if it is a repository (default last commit)")
	testCmd.Flags().StringVarP(&clusterExecName, "cluster", "c", "", "cluster name on which to run the project (default cluster)")
	testCmd.Flags().BoolVarP(&watchFlag, "watch", "w", false, "watch test executions with results (default false)")
}
