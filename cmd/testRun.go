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
	runCmd          = &cobra.Command{
		Use:   "run",
		Short: "The test run command allows you to perform the flakiness analysis of a project",
		Long: `The test run command allows you to carry out the flakiness analysis of a local or remote project,
		correctly configured with a .flaky.yaml file, by sending it to one of the added clusters.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Test Run Called")
		},
	}
)

func init() {
	testCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&sourceUrl, "source", "s", ".", "url at which the project can be found, local path or repository url")
	runCmd.Flags().StringVarP(&commitId, "commitId", "i", "", "hash of the commit to refer to if it is a repository (default last commit)")
	runCmd.Flags().StringVarP(&clusterExecName, "cluster", "c", "", "cluster name on which to run the project (default cluster)")
	runCmd.Flags().BoolVarP(&watchFlag, "watch", "w", false, "watch test executions with results (default false)")
}
