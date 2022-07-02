/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	setCmd = &cobra.Command{
		Use:   "set [name of the cluster]",
		Short: "The cluster set command allows you to set a new default cluster",
		Long:  `The cluster set command allows you to set a new default cluster from the list of those you have access to.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newDefaultClusterName := args[0]

			var clusters []map[string]interface{}
			getClusters(&clusters)

			clusterIndex := isClusterPresent(&clusters, newDefaultClusterName)

			if clusterIndex == -1 {
				fmt.Printf("Error: Cluster %s not found.\n", newDefaultClusterName)
				os.Exit(0)
			}

			viper.Set("defaultCluster", newDefaultClusterName)
			viper.WriteConfig()
			fmt.Printf("Cluster %s is the new default cluster.\n", newDefaultClusterName)
		},
	}
)

func init() {
	clusterCmd.AddCommand(setCmd)
}
