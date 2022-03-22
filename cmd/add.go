/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	clusterName    string
	clusterUrl     string
	userEmail      string
	userPassword   string
	defaultCluster bool
	addCmd         = &cobra.Command{
		Use:   "add",
		Short: "The cluster add command allows you to add a new cluster",
		Long:  `The cluster add command allows you to authenticate to a new cluster and add it to the list of usable clusters.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cluster Add Called")
		},
	}
)

func init() {
	clusterCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&clusterName, "name", "n", "", "name that uniquely identifies a cluster (required)")
	addCmd.MarkFlagRequired("name")
	addCmd.Flags().StringVarP(&clusterUrl, "url", "u", "", "url at which the project can be found, local path or repository url (required)")
	addCmd.MarkFlagRequired("url")
	addCmd.Flags().StringVarP(&userEmail, "email", "e", "", "email of the account with which you are registered to the cluster (required)")
	addCmd.MarkFlagRequired("email")
	addCmd.Flags().StringVarP(&userPassword, "password", "p", "", "password of the account with which you are registered to the cluster (required)")
	addCmd.MarkFlagRequired("password")
	addCmd.Flags().BoolVarP(&defaultCluster, "default", "d", false, "set cluster as default (default false)")
}
