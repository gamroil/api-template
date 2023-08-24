package config

import (
	"log"

	"github.com/spf13/viper"
)

// Configuration is the struct representation of the config yaml. Instantiation should occur through the Configure
// function as it creates internal resources.
type Configuration struct {
	Environment string `mapstructure:"environment"`
	Port        int    `mapstructure:"port"`
}

// Configure is the intended way to instantiate a Configuration. This method should be used over direct instantiation
// because it creates internal resources.
func Configure(filepath string) *Configuration {
	viper.SetConfigFile(filepath)
	viper.SetEnvPrefix("OVERRIDE")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("reading config: ", err)
	}

	var conf *Configuration

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalln("unmarshaling config: ", err)
	}

	return conf
}
