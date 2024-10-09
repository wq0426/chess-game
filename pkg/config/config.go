package config

import (
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"os"
)

var ConfigInstance *viper.Viper
var Rdb *redis.Client

func NewConfig() *viper.Viper {
	if ConfigInstance == nil {
		envConf := os.Getenv("APP_CONF")
		if envConf == "" {
			flag.StringVar(&envConf, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
			flag.Parse()
		}
		if envConf == "" {
			envConf = "config/local.yml"
		}
		fmt.Println("load conf file:", envConf)
		ConfigInstance = getConfig(envConf)
	}
	if Rdb == nil {
		Rdb = getRdb()
	}
	return ConfigInstance
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

func getRdb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     ConfigInstance.GetString("data.redis.addr"),
		Password: ConfigInstance.GetString("data.redis.password"), // no password set
		DB:       ConfigInstance.GetInt("data.redis.db"),          // use default DB
	})
}
