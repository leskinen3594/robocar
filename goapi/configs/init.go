package configs

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
