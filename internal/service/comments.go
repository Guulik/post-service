package service

import (
	"context"
	"posts/internal/constants"
	"posts/internal/model"
	"posts/internal/repository"
)

type CommentsService struct {
	repo       repository.Comments
	PostGetter PostProvider
}

type PostProvider interface {
	GetPostById(ctx context.Context, id int) (model.Post, error)
}

func NewCommentsService(repo repository.Comments, getter PostProvider) *CommentsService {
	return &CommentsService{repo: repo, PostGetter: getter}
}

func (c CommentsService) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	if len(comment.Content) >= constants.MaxContentLength {
		//TODO: handle error
	}

	post, err := c.PostGetter.GetPostById(ctx, comment.Post)
	if err != nil {
		//TODO: handle error
	}

	if !post.CommentsAllowed {
		//TODO: handle error
	}

	newComment, err := c.repo.CreateComment(ctx, comment)
	if err != nil {
		//TODO: handle error
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(ctx context.Context, postId int, limit int, offset int) ([]*model.Comment, error) {

	if postId <= 0 {
		//TODO: handle error
	}
	comments, err := c.repo.GetCommentsByPost(ctx, postId, limit, offset)
	if err != nil {
		//TODO: handle error
	}

	return comments, nil
}

func (c CommentsService) GetRepliesOfComment(ctx context.Context, commentId int) ([]*model.Comment, error) {

	if commentId <= 0 {
		//TODO: handle error
	}

	comments, err := c.repo.GetRepliesOfComment(ctx, commentId)
	if err != nil {
		//TODO: handle error
	}

	return comments, nil

}
