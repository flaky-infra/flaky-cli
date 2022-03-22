/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "The cluster add command allows you to add a new cluster",
	Long:  `The cluster add command allows you to authenticate to a new cluster and add it to the list of usable clusters.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cluster Add Called")
	},
}

func init() {
	clusterCmd.AddCommand(addCmd)
}
