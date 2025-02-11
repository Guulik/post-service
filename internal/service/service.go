package service

import (
	"context"
	"posts/internal/model"
	"posts/internal/repository"
)

type Services struct {
	Posts
	Comments
}

func NewServices(repo *repository.Repo) *Services {
	return &Services{
		Posts:    NewPostsService(repo.Posts),
		Comments: NewCommentsService(repo.Comments, repo.Posts),
	}
}

type Posts interface {
	CreatePost(ctx context.Context, post model.Post) (model.Post, error)
	GetPostById(ctx context.Context, id int) (model.Post, error)
	GetAllPosts(ctx context.Context, limit int, offset int) ([]*model.Post, error)
}

type Comments interface {
	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	GetCommentsByPost(ctx context.Context, postId int, limit int, offset int) ([]*model.Comment, error)
	GetRepliesOfComment(ctx context.Context, commentId int) ([]*model.Comment, error)
}
