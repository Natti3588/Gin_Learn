package handler

import (
	"net/http"

	"backend/internal/dto"

	"github.com/gin-gonic/gin"
)

// @Summary  チャットの追加と表示
// @Tags     chat
// @Accept   json
// @Produce  json
// @Param    request body dto.ChatRequest true "チャットリクエスト"
// @Success  200 {object} dto.ChatResponse
// @Router   /posts [post]
func ChatHandler(c *gin.Context) {
	var req dto.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.ChatResponse{
		Reply: "受け取ったメッセージ: " + req.Message,
	}
	c.JSON(http.StatusOK, res)
}
