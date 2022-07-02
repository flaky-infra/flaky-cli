/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [name of the cluster]",
	Short: "The cluster set command allows you to set a new default cluster",
	Long:  `The cluster set command allows you to set a new default cluster from the list of those you have access to.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster Set Called")
	},
}

func init() {
	clusterCmd.AddCommand(setCmd)
}