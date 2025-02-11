package in_memory

import (
	"context"
	"posts/internal/model"

	"sync"
	"time"
)

type CommentsInMemory struct {
	mu       sync.RWMutex
	comments []model.Comment
}

func NewCommentsInMemory(count int) *CommentsInMemory {
	return &CommentsInMemory{
		comments: make([]model.Comment, 0, count),
	}
}

func (c *CommentsInMemory) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	t := time.Now()

	comment.ID = len(c.comments) + 1
	comment.CreatedAt = t

	c.comments = append(c.comments, comment)

	return comment, nil

}

func (c *CommentsInMemory) GetCommentsByPost(ctx context.Context, postId, limit, offset int) ([]*model.Comment, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	var res []*model.Comment

	for _, comment := range c.comments {
		if comment.ReplyTo == nil && comment.Post == postId {
			res = append(res, &comment)
		}
	}

	if offset > len(res) {
		return nil, nil
	}

	if offset+limit > len(res) || limit == -1 {
		return res[offset:], nil
	}

	return res[offset : offset+limit], nil
}

func (c *CommentsInMemory) GetRepliesOfComment(ctx context.Context, commentId int) ([]*model.Comment, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if commentId > len(c.comments) {
		return nil, nil
	}

	var res []*model.Comment

	for _, comment := range c.comments {
		if comment.ReplyTo != nil && *comment.ReplyTo == commentId {
			res = append(res, &comment)
		}
	}

	return res, nil
}
