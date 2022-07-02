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

var rmCmd = &cobra.Command{
	Use:   "rm [name of the cluster]",
	Short: "The cluster rm command allows you to delete a cluster from the list of those to which you have access",
	Long:  `The cluster rm command allows you to delete a cluster from the list of those to which you have access`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteClusterName := args[0]

		var clusters []map[string]interface{}
		getClusters(&clusters)

		var newClusters []map[string]interface{}

		clusterIndex := isClusterPresent(&clusters, deleteClusterName)

		if clusterIndex != -1 {
			newClusters = append(clusters[:clusterIndex], clusters[clusterIndex+1:]...)
		} else {
			fmt.Printf("Error: Cluster %s not found.\n", deleteClusterName)
			os.Exit(0)
		}

		fmt.Printf("Cluster %s has been removed.\n", deleteClusterName)

		if viper.Get("defaultCluster") == deleteClusterName {
			viper.Set("defaultCluster", newClusters[0]["clusterName"])
			fmt.Printf("Since %s has been dropped, the new default cluster is %s.\n", deleteClusterName, newClusters[0]["clusterName"])
		}

		viper.Set("clusters", newClusters)
		viper.WriteConfig()
	},
}

func init() {
	clusterCmd.AddCommand(rmCmd)
}
