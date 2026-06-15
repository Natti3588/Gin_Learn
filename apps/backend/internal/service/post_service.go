package service

import (
	"errors"
	"strings"

	"backend/internal/model"
)

// ErrTitleRequired はタイトル未入力時に返すバリデーションエラー。
var ErrTitleRequired = errors.New("title is required")

// PostRepository は service が依存する DB 操作の窓口（インターフェース）。
// 具象は repository パッケージ、テストでは fake を差し込む。
type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
}

// PostService は投稿に関する業務ロジックを担う。
type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{repo: repo}
}

// Create はタイトル・本文を整形・検証して投稿を保存する。
func (s *PostService) Create(title, body string) (*model.Post, error) {
	title = strings.TrimSpace(title)
	body = strings.TrimSpace(body)
	if title == "" {
		return nil, ErrTitleRequired
	}

	post := &model.Post{Title: title, Body: body}
	if err := s.repo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

// FindAll は保存されている投稿をすべて取得する。
func (s *PostService) FindAll() ([]model.Post, error) {
	return s.repo.FindAll()
}
