package in_memory

import (
	"database/sql"
	"posts/internal/model"
	"sync"
	"time"
)

type PostsInMemory struct {
	mu    sync.RWMutex
	posts []*model.Post
}

func NewPostsInMemory(count int) *PostsInMemory {
	return &PostsInMemory{

		posts: make([]*model.Post, 0, count),
	}
}

func (p *PostsInMemory) CreatePost(post model.Post) (model.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	post.ID = len(p.posts) + 1
	post.CreatedAt = time.Now()

	p.posts = append(p.posts, &post)

	return post, nil
}

func (p *PostsInMemory) GetPostById(id int) (model.Post, error) {

	p.mu.RLock()
	defer p.mu.RUnlock()

	if id > len(p.posts) || id <= 0 {
		return model.Post{}, sql.ErrNoRows
	}

	return *p.posts[id-1], nil
}

func (p *PostsInMemory) GetAllPosts(limit, offset int) ([]*model.Post, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if offset > len(p.posts) {
		return nil, nil
	}

	if offset+limit > len(p.posts) || limit == -1 {
		return p.posts[offset:], nil
	}

	return p.posts[offset : offset+limit], nil
}
