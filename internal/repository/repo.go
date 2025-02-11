package repository

import (
	"posts/internal/model"
)

type Repo struct {
	RepoPosts
	RepoComments
}

func NewRepo(posts RepoPosts, comments RepoComments) *Repo {
	return &Repo{
		RepoPosts:    posts,
		RepoComments: comments,
	}
}

type RepoPosts interface {
	CreatePost(post model.Post) (model.Post, error)
	GetPostById(id int) (model.Post, error)
	GetAllPosts(limit, offset int) ([]*model.Post, error)
}

type RepoComments interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetCommentsByPost(postId, limit, offset int) ([]*model.Comment, error)
	GetRepliesOfComment(commentId int) ([]*model.Comment, error)
}
