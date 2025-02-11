package model

import "time"

type Post struct {
	ID              int       `json:"id" db:"id"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	Name            string    `json:"name" db:"name"`
	Content         string    `json:"content" db:"content"`
	CommentsAllowed bool      `json:"commentsAllowed" db:"comments_allowed"`
}

func (p InputPost) FromInput() Post {
	return Post{
		Name:            p.Name,
		Content:         p.Content,
		CommentsAllowed: p.CommentsAllowed,
	}
}
