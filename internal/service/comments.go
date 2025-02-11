package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"posts/internal/constants"
	"posts/internal/model"
	"posts/internal/repository"
)

type CommentsService struct {
	repo       repository.RepoComments
	PostGetter PostProvider
	log        *slog.Logger
}

type PostProvider interface {
	GetPostById(id int) (model.Post, error)
}

func NewCommentsService(repo repository.RepoComments, getter PostProvider, log *slog.Logger) *CommentsService {
	return &CommentsService{repo: repo,
		PostGetter: getter,
		log:        log,
	}
}

func (c CommentsService) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	const op = "Service.CreateComment"
	log := c.log.With(
		slog.String("op", op),
	)

	post, err := c.PostGetter.GetPostById(comment.Post)
	if err != nil {
		log.Error(constants.PostNotFoundError, err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return model.Comment{}, fmt.Errorf("%s: %w", constants.PostNotFoundError, err)
		}
	}

	if !post.CommentsAllowed {
		log.Error(constants.CommentsNotAllowedError)
		return model.Comment{}, fmt.Errorf("%s", constants.CommentsNotAllowedError)
	}

	newComment, err := c.repo.CreateComment(comment)
	if err != nil {
		log.Error(constants.CreatingCommentError, err.Error())
		return model.Comment{}, fmt.Errorf("%s: %w", constants.CreatingCommentError, err)
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(ctx context.Context, postId int, limit int, offset int) ([]*model.Comment, error) {
	const op = "Service.GetCommentsByPost"
	log := c.log.With(
		slog.String("op", op),
	)

	comments, err := c.repo.GetCommentsByPost(postId, limit, offset)
	if err != nil {
		log.Error(constants.GettingCommentError)
		return nil, fmt.Errorf("%s, postId=%d: %w", constants.GettingCommentError, postId, err)
	}

	return comments, nil
}

func (c CommentsService) GetRepliesOfComment(ctx context.Context, commentId int) ([]*model.Comment, error) {
	const op = "Service.GetRepliesOfComment"
	log := c.log.With(
		slog.String("op", op),
	)

	comments, err := c.repo.GetRepliesOfComment(commentId)
	if err != nil {
		log.Error(constants.GettingRepliesError, commentId)
		return nil, fmt.Errorf("%s, commentId=%d: %w", constants.GettingRepliesError, commentId, err)
	}

	return comments, nil

}
