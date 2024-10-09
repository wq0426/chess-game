package server

import (
	"chess/internal/handler"
	"chess/pkg/log"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	logger *log.Logger,
	chessHandler *handler.RankHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	noAuthGroup := r.Group("/api/v1")
	// 查询排行榜列表
	noAuthGroup.GET("/rank", chessHandler.GetRank)
	// 推动用户通关记录
	noAuthGroup.POST("/rank", chessHandler.PostRank)

	return r
}
