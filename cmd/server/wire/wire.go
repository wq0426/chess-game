//go:build wireinject
// +build wireinject

package wire

import (
	"chess/internal/handler"
	"chess/internal/repository"
	"chess/internal/server"
	"chess/internal/service"
	"chess/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewRankRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewRankService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewRankHandler,
)

func NewWire(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
