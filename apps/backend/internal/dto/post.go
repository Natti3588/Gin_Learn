package dto

// CreatePostRequest は投稿作成エンドポイントのリクエストボディ。
type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PostResponse は投稿のレスポンスボディ。
type PostResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
