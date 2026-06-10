package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary  チャットの追加と表示
// @Tags     chat
// @Accept   json
// @Produce  json
// @Param    request body models.ChatRequest true "チャットリクエスト"
// @Success  200 {object} models.ChatResponse
// @Router   /posts [post]
func ChatHandler(c *gin.Context) {
	var req models.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := models.ChatResponse {
		Reply: "受け取ったメッセージ: " + req.Message,
	}
	c.JSON(http.StatusOK, res)
}