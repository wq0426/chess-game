package handler

import (
	"chess/internal/model"
	"chess/internal/service"
	"chess/pkg/helper/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type RankHandler struct {
	*Handler
	rankService service.RankService
}

func NewRankHandler(
	handler *Handler,
	rankService service.RankService,
) *RankHandler {
	return &RankHandler{
		Handler:     handler,
		rankService: rankService,
	}
}

func (h *RankHandler) GetRank(ctx *gin.Context) {
	rankList, err := h.rankService.GetRank(ctx)
	if err != nil {
		return
	}
	resp.HandleSuccess(ctx, rankList)
}

func (h *RankHandler) PostRank(ctx *gin.Context) {
	params := &model.UserRankParam{}
	if err := ctx.ShouldBindJSON(params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	// 验证参数是否有效
	if !params.Validate() {
		resp.HandleError(ctx, http.StatusBadRequest, 1, "参数格式错误", nil)
		return
	}
	_, err := h.rankService.AddUserRank(ctx, params)
	h.logger.Info("Add rank", zap.Any("param", params))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, "添加失败", nil)
		return
	}
	resp.HandleSuccess(ctx, true)
}
