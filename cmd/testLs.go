/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/flaky-infra/flaky-cli/util"
	"github.com/spf13/cobra"
)

var lsTestCmd = &cobra.Command{
	Use:   "ls",
	Short: "The cluster ls command allows you to view the list of last executed tests",
	Long:  `The cluster ls command allows you to view the list of last executed tests.`,
	Run: func(cmd *cobra.Command, args []string) {
		var testRuns []map[string]interface{}
		util.GetSliceMapViper("testRuns", &testRuns)

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "TEST RUN ID\tDATE\tCLUSTER\t")
		for _, testRun := range testRuns {
			fmt.Fprintf(w, "%s\t%s\t%s\t\n", testRun["testRunId"], testRun["date"], testRun["cluster"])
		}
		w.Flush()
	},
}

func init() {
	testCmd.AddCommand(lsTestCmd)
}
