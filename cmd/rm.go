/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "The cluster rm command allows you to delete a cluster from the list of those to which you have access",
	Long:  `The cluster rm command allows you to delete a cluster from the list of those to which you have access`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster Rm Called")
	},
}

func init() {
	clusterCmd.AddCommand(rmCmd)
}