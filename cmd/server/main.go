package main

import (
	"chess/cmd/server/wire"
	"chess/pkg/config"
	"chess/pkg/http"
	"chess/pkg/log"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", conf.GetString("http.host")+":"+conf.GetString("http.port")))

	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
