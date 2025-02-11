package model

import "time"

type Comment struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Content   string    `json:"content" db:"content"`
	Post      int       `json:"post" db:"post"`
	ReplyTo   *int      `json:"replyTo,omitempty" db:"reply_to"`
}

func (c InputComment) FromInput() Comment {
	return Comment{
		Content: c.Content,
		Post:    c.Post,
		ReplyTo: c.ReplyTo,
	}
}
