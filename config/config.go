package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	homeDir, _ := os.UserHomeDir()
	viper.AddConfigPath(homeDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".drippypale-todo-list")

	if err := viper.ReadInConfig(); err != nil {
		defaultConfig()
	}
}

func defaultConfig() {
	homeDir, _ := os.UserHomeDir()

	// list of default config vars
	viper.Set("csvPath", fmt.Sprintf("%v/%v", homeDir, "drippypale-todo-list.csv"))
	viper.Set("timeFormat", "2006-01-02 15:04")
	// end of list

	err := viper.SafeWriteConfig()
	if err != nil {
		panic("Can not read the configuration file.")
	}
}
