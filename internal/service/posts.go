package service

import (
	"context"
	"posts/internal/constants"
	"posts/internal/model"
	"posts/internal/repository"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{repo: repo}
}

func (p PostsService) CreatePost(ctx context.Context, post model.Post) (model.Post, error) {

	if len(post.Content) >= constants.MaxContentLength {
		//TODO: handle error
	}

	newPost, err := p.repo.CreatePost(ctx, post)
	if err != nil {
		//TODO: handle error
	}

	return newPost, nil

}

func (p PostsService) GetPostById(ctx context.Context, postId int) (model.Post, error) {
	if postId <= 0 {
		//TODO: handle error
	}

	post, err := p.repo.GetPostById(ctx, postId)
	if err != nil {
		//TODO: handle error
	}

	return post, nil
}

func (p PostsService) GetAllPosts(ctx context.Context, limit, offset int) ([]*model.Post, error) {

	posts, err := p.repo.GetAllPosts(ctx, limit, offset)
	if err != nil {
		//TODO: handle error
	}

	return posts, nil
}
