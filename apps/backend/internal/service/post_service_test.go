package service

import (
	"errors"
	"testing"

	"backend/internal/model"
)

// fakePostRepo は DB を使わないテスト用の PostRepository 実装。
type fakePostRepo struct {
	created  []*model.Post
	findAll  []model.Post
	err      error
}

func (f *fakePostRepo) Create(post *model.Post) error {
	if f.err != nil {
		return f.err
	}
	f.created = append(f.created, post)
	return nil
}

func (f *fakePostRepo) FindAll() ([]model.Post, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.findAll, nil
}

func TestCreate_TrimsTitleAndBodyThenSaves(t *testing.T) {
	repo := &fakePostRepo{}
	svc := NewPostService(repo)

	post, err := svc.Create("  Hello  ", "  world  ")
	if err != nil {
		t.Fatalf("想定外のエラー: %v", err)
	}
	if post.Title != "Hello" || post.Body != "world" {
		t.Fatalf("trim されていない: title=%q body=%q", post.Title, post.Body)
	}
	if len(repo.created) != 1 {
		t.Fatalf("repository に1件保存されるはず: %d件", len(repo.created))
	}
	if repo.created[0].Title != "Hello" {
		t.Fatalf("保存値が trim されていない: %q", repo.created[0].Title)
	}
}

func TestCreate_RejectsEmptyTitle(t *testing.T) {
	repo := &fakePostRepo{}
	svc := NewPostService(repo)

	_, err := svc.Create("   ", "body")
	if !errors.Is(err, ErrTitleRequired) {
		t.Fatalf("ErrTitleRequired を期待: %v", err)
	}
	if len(repo.created) != 0 {
		t.Fatalf("バリデーション失敗時は repository を呼ばないはず: %d件", len(repo.created))
	}
}

func TestFindAll_ReturnsAllPosts(t *testing.T) {
	repo := &fakePostRepo{findAll: []model.Post{
		{Title: "A", Body: "a"},
		{Title: "B", Body: "b"},
	}}
	svc := NewPostService(repo)

	posts, err := svc.FindAll()
	if err != nil {
		t.Fatalf("想定外のエラー: %v", err)
	}
	if len(posts) != 2 {
		t.Fatalf("2件返るはず: %d件", len(posts))
	}
	if posts[0].Title != "A" || posts[1].Title != "B" {
		t.Fatalf("内容が一致しない: %+v", posts)
	}
}
