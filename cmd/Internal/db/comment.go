package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kanishkanaik/go-rest-api-course/cmd/Internal/comment"
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
