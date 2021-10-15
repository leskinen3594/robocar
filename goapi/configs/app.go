package configs

import "github.com/spf13/viper"

func AppConfig() (string, int) {
	host := viper.GetString("app.host")
	port := viper.GetInt("app.port")

	return host, port
}
