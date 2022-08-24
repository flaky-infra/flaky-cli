package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetSliceMapViper(key string, out *[]map[string]interface{}) *[]map[string]interface{} {
	err := viper.UnmarshalKey(key, &out)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}

	return out
}
