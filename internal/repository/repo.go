package repository

import (
	"context"
	"posts/internal/model"
)

type Repo struct {
	Posts
	Comments
}

func NewRepo(posts Posts, comments Comments) *Repo {
	return &Repo{
		Posts:    posts,
		Comments: comments,
	}
}

type Posts interface {
	CreatePost(ctx context.Context, post model.Post) (model.Post, error)
	GetPostById(ctx context.Context, id int) (model.Post, error)
	GetAllPosts(ctx context.Context, limit, offset int) ([]*model.Post, error)
}

type Comments interface {
	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	GetCommentsByPost(ctx context.Context, postId, limit, offset int) ([]*model.Comment, error)
	GetRepliesOfComment(ctx context.Context, commentId int) ([]*model.Comment, error)
}
