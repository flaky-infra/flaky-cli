/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "The cluster set command allows you to set a new default cluster",
	Long:  `The cluster set command allows you to set a new default cluster from the list of those you have access to.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster Set Called")
	},
}

func init() {
	clusterCmd.AddCommand(setCmd)
}
