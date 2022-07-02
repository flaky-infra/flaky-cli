/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getClusters(clusters *[]map[string]interface{}) *[]map[string]interface{} {
	err := viper.UnmarshalKey("clusters", &clusters)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return clusters
}

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "The cluster command can be used for cluster management",
	Long: `The cluster command can be used for cluster management, you can:
add a cluster, set a cluster as default, remove a cluster and view the added clusters.`,
}

func init() {
	rootCmd.AddCommand(clusterCmd)
}
