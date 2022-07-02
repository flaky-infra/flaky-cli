/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lsClusterCmd = &cobra.Command{
	Use:   "ls",
	Short: "The cluster ls command allows you to view the list of your clusters",
	Long:  `The cluster ls command allows you to view the list of clusters to which you have access.`,
	Run: func(cmd *cobra.Command, args []string) {
		var clusters []map[string]interface{}
		getClusters(&clusters)
		defaultCluster := viper.Get("defaultCluster")

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintln(w, "NAME\tURL\tUSERNAME\tDEFAULT\t")
		for _, cluster := range clusters {
			if cluster["clusterName"] == defaultCluster {
				fmt.Fprintf(w, "%s\t%s\t%s\tx\t\n", cluster["clusterName"], cluster["clusterUrl"], cluster["userEmail"])
			} else {
				fmt.Fprintf(w, "%s\t%s\t%s\t\t\n", cluster["clusterName"], cluster["clusterUrl"], cluster["userEmail"])
			}
		}
		w.Flush()

	},
}

func init() {
	clusterCmd.AddCommand(lsClusterCmd)
}
