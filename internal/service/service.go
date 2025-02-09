package service

import (
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
	CreatePost(post model.Post) (model.Post, error)
	GetPostById(id int) (model.Post, error)
	GetAllPosts(page, pageSize *int) ([]model.Post, error)
}

type Comments interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetCommentsByPost(postId int, page *int, pageSize *int) ([]*model.Comment, error)
	GetRepliesOfComment(commentId int, depth int32) ([]*model.Comment, error)
}
