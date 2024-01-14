package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kanishkanaik/go-rest-api-course/cmd/Internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Body:   c.Body.String,
		Author: c.Author.String,
		Slug:   c.Slug.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {

	var cmtRow CommentRow
	row := d.Client.QueryRowxContext(ctx,
		`SELECT id, slug, body, author
	FROM comments
	WHERE id = $1`, uuid)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("Error Fecthing Comment By ID %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(ctx,
		`INSERT INTO comments (id, slug, author, body) VALUES (:id, :slug, :author, :body)`, postRow)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("Cound not Add Comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("Could not Close Row: %w", err)
	}

	return cmt, nil

}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(ctx, `DELETE FROM comments where id $1`, id)
	if err != nil {
		return fmt.Errorf("Could not able to Delete Comment")
	}

	return nil
}

func (d *Database) UpdateComment(ctx context.Context, cmt comment.Comment, id string) (comment.Comment, error) {

	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(ctx,
		`UPDATE comments SET 
		slug = :slug, 
		author= :author,
		body: =body
		WHERE id= :id`, cmtRow)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("Cound not Update Comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("Could not Close Row: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil

}
