package service

import (
	"posts/internal/constants"
	"posts/internal/lib/pagination"
	"posts/internal/model"
	"posts/internal/repository"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{repo: repo}
}

func (p PostsService) CreatePost(post model.Post) (model.Post, error) {

	if len(post.Content) >= constants.MaxContentLength {
		//TODO: handle error
	}

	newPost, err := p.repo.CreatePost(post)
	if err != nil {
		//TODO: handle error
	}

	return newPost, nil

}

func (p PostsService) GetPostById(postId int) (model.Post, error) {

	if postId <= 0 {
		//TODO: handle error
	}

	post, err := p.repo.GetPostById(postId)
	if err != nil {
		//TODO: handle error
	}

	return post, nil
}

func (p PostsService) GetAllPosts(page, pageSize *int) ([]model.Post, error) {

	if page != nil && *page < 0 {
		//TODO: handle error
	}

	if pageSize != nil && *pageSize < 0 {
		//TODO: handle error
	}

	offset, limit := pagination.GetOffsetAndLimit(page, pageSize)

	posts, err := p.repo.GetAllPosts(limit, offset)
	if err != nil {
		//TODO: handle error
	}

	return posts, nil
}
