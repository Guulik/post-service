package service

import (
	"posts/internal/constants"
	"posts/internal/lib/pagination"
	"posts/internal/model"
	"posts/internal/repository"
)

type CommentsService struct {
	repo       repository.Comments
	PostGetter PostGetter
}

type PostGetter interface {
	GetPostById(id int) (model.Post, error)
}

func NewCommentsService(repo repository.Comments, getter PostGetter) *CommentsService {
	return &CommentsService{repo: repo, PostGetter: getter}
}

func (c CommentsService) CreateComment(comment model.Comment) (model.Comment, error) {
	if len(comment.Content) >= constants.MaxContentLength {
		//TODO: handle error
	}

	post, err := c.PostGetter.GetPostById(comment.Post)
	if err != nil {
		//TODO: handle error
	}

	if !post.CommentsAllowed {
		//TODO: handle error
	}

	newComment, err := c.repo.CreateComment(comment)
	if err != nil {
		//TODO: handle error
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*model.Comment, error) {

	if postId <= 0 {
		//TODO: handle error
	}

	if page != nil && *page < 0 {
		//TODO: handle error
	}

	if pageSize != nil && *pageSize < 0 {
		//TODO: handle error
	}

	offset, limit := pagination.GetOffsetAndLimit(page, pageSize)

	comments, err := c.repo.GetCommentsByPost(postId, limit, offset)
	if err != nil {
		//TODO: handle error
	}

	return comments, nil
}

func (c CommentsService) GetRepliesOfComment(commentId int, depth int32) ([]*model.Comment, error) {

	if commentId <= 0 {
		//TODO: handle error
	}

	comments, err := c.repo.GetRepliesOfComment(commentId, depth)
	if err != nil {
		//TODO: handle error
	}

	return comments, nil

}
