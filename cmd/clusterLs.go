/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsClusterCmd = &cobra.Command{
	Use:   "ls",
	Short: "The cluster ls command allows you to view the list of your clusters",
	Long:  `The cluster ls command allows you to view the list of clusters to which you have access.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster Ls Called")
	},
}

func init() {
	clusterCmd.AddCommand(lsClusterCmd)
}
