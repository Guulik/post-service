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

type PostsService struct {
	repo repository.RepoPosts
	log  *slog.Logger
}

func NewPostsService(repo repository.RepoPosts, log *slog.Logger) *PostsService {
	return &PostsService{repo: repo, log: log}
}

func (p PostsService) CreatePost(ctx context.Context, post model.Post) (model.Post, error) {
	const op = "Service.CreatePost"
	log := p.log.With(
		slog.String("op", op),
	)

	newPost, err := p.repo.CreatePost(post)
	if err != nil {
		log.Error(constants.CreatingPostError, slog.String("err:", err.Error()))
		return model.Post{}, fmt.Errorf("%s: %w", constants.CreatingPostError, err)
	}

	return newPost, nil

}

func (p PostsService) GetPostById(ctx context.Context, postId int) (model.Post, error) {
	const op = "Service.GetPostById"
	log := p.log.With(
		slog.String("op", op),
	)

	post, err := p.repo.GetPostById(postId)
	if err != nil {
		log.Error(constants.GetPostError, slog.String("err:", err.Error()))
		if errors.Is(err, sql.ErrNoRows) {
			return model.Post{}, fmt.Errorf("%s, postId=%d: %w", constants.PostNotFoundError, postId, err)
		}
		return model.Post{}, fmt.Errorf("%s, postId=%d: %w", constants.GetPostError, postId, err)
	}

	return post, nil
}

func (p PostsService) GetAllPosts(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	const op = "Service.GetAllPosts"
	log := p.log.With(
		slog.String("op", op),
	)

	posts, err := p.repo.GetAllPosts(limit, offset)
	if err != nil {
		log.Error(constants.GetPostError, slog.String("err:", err.Error()))
		return nil, fmt.Errorf("%s: %w", constants.GetPostError, err)
	}

	return posts, nil
}
