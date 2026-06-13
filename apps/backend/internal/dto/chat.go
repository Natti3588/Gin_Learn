package dto

// ChatRequest はチャットエンドポイントへのリクエストボディ。
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

// ChatResponse はチャットエンドポイントのレスポンスボディ。
type ChatResponse struct {
	Reply string `json:"reply"`
}
