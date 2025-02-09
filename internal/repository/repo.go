package repository

import "posts/internal/model"

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
	CreatePost(post model.Post) (model.Post, error)
	GetPostById(id int) (model.Post, error)
	GetAllPosts(limit, offset int) ([]model.Post, error)
}

type Comments interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetCommentsByPost(postId, limit, offset int) ([]*model.Comment, error)
	GetRepliesOfComment(commentId int, depth int32) ([]*model.Comment, error)
}
