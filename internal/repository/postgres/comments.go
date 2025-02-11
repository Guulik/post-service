package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"posts/internal/model"
)

type CommentsPostgres struct {
	db *sqlx.DB
}

func NewCommentsPostgres(db *sqlx.DB) *CommentsPostgres {
	return &CommentsPostgres{db: db}
}

func (c CommentsPostgres) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {

	tx, err := c.db.Begin()
	if err != nil {
		return model.Comment{}, err
	}

	query := `INSERT INTO comments (content, post, reply_to) 
				VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	row := tx.QueryRow(query, comment.Content, comment.Post, comment.ReplyTo)
	if err := row.Scan(&comment.ID, &comment.CreatedAt); err != nil {
		tx.Rollback()
		return model.Comment{}, err
	}

	return comment, tx.Commit()

}

func (c CommentsPostgres) GetCommentsByPost(ctx context.Context, postId, limit, offset int) ([]*model.Comment, error) {

	query := `SELECT * FROM comments 
         WHERE post = $1 AND reply_to IS NULL 
         ORDER BY created_at 
         OFFSET $2`

	args := []interface{}{postId, offset}

	if limit >= 0 {
		query += " LIMIT $3"
		args = append(args, limit)
	}

	var comments []*model.Comment

	if err := c.db.Select(&comments, query, args...); err != nil {
		return nil, err
	}

	return comments, nil
}

func (c CommentsPostgres) GetRepliesOfComment(ctx context.Context, commentId int) ([]*model.Comment, error) {
	query := `SELECT * FROM comments WHERE reply_to = $1`

	var comments []*model.Comment

	if err := c.db.Select(&comments, query, commentId); err != nil {
		return nil, err
	}

	return comments, nil
}
