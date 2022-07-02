/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "flaky-cli",
	Short: "This CLI allows to obtain information related to the flaky tests of an application",
	Long: `This CLI allows to obtain information related to the flaky tests of an application.
You need to add a cluster that allows you to run tests based on your
project configuration (specified in a .flaky.yaml file).`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("An error has occurred")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	flakyPath := filepath.Join(home, ".flaky")
	viper.AddConfigPath(flakyPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		os.Mkdir(flakyPath, 0755)
		err = viper.SafeWriteConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
