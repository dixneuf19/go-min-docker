package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Init initializes config
func Init() {
	// Get Config from yml files
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("%s: %s", "Unable to read Config Files in config/*", err))
	}

	if _, err := os.Stat("./secret.yml"); !os.IsNotExist(err) {
		viper.SetConfigName("secret")
		viper.AddConfigPath("./")
		err := viper.MergeInConfig()
		if err != nil {
			panic(fmt.Sprintf("%s: %s", "Unable to read Config File secret.yml", err))
		}
	} else {
		fmt.Println("Secret file not found")
	}

	// Overrides configs with ENV variables, ex: grpc.port with GRPC_PORT
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	fmt.Printf("Settings are %#v \n", viper.AllSettings())
}

// GetString returns a string config from name
func GetString(key string) string {
	return viper.GetString(key)
}
