package handler

import (
	"errors"
	"net/http"

	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

// PostHandler は投稿系エンドポイントの HTTP 入出力を担う。
type PostHandler struct {
	svc *service.PostService
}

func NewPostHandler(svc *service.PostService) *PostHandler {
	return &PostHandler{svc: svc}
}

// Create godoc
// @Summary  投稿の作成
// @Tags     posts
// @Accept   json
// @Produce  json
// @Param    request body dto.CreatePostRequest true "投稿作成リクエスト"
// @Success  201 {object} dto.PostResponse
// @Failure  400 {object} map[string]string
// @Failure  500 {object} map[string]string
// @Router   /posts [post]
func (h *PostHandler) Create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.svc.Create(req.Title, req.Body)
	if err != nil {
		// 業務ルール違反は 400、それ以外は 500 に振り分ける
		if errors.Is(err, service.ErrTitleRequired) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.PostResponse{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	})
}

// List godoc
// @Summary  投稿の全件取得
// @Tags     posts
// @Produce  json
// @Success  200 {array} dto.PostResponse
// @Failure  500 {object} map[string]string
// @Router   /posts [get]
func (h *PostHandler) List(c *gin.Context) {
	posts, err := h.svc.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]dto.PostResponse, 0, len(posts))
	for _, p := range posts {
		res = append(res, dto.PostResponse{
			ID:    p.ID,
			Title: p.Title,
			Body:  p.Body,
		})
	}
	c.JSON(http.StatusOK, res)
}
