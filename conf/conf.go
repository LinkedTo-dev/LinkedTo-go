package conf

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var LoginInfo map[string]string

func init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	LoginInfo = viper.GetStringMapString("sql")
}
