package handler

import (
	"chess/pkg/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewHandler(logger *log.Logger, db *gorm.DB) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

func (h *Handler) StartTrans(ctx *gin.Context) (*gorm.DB, error) {
	tx := h.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	// 方便数据库统一回滚
	ctx.Set("tx_key", "bind")
	ctx.Set("tx_key_bind", tx)
	return tx, nil
}
