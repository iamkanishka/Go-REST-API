package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("Failed to Fetch comment by Id")
	ErrorCreatingComment = errors.New("Failed to Create comment by Id")
	ErrorUpdatingComment = errors.New("Failed to Update comment by Id")
	ErrorNotImplemented  = errors.New("Method not Implemented")
)

// Comment - Representation Structure of Comment fo
// Service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store -  this interface defines all the
// methods that out service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, Comment, string) (Comment, error)
}

// Service - a struct on whcih all our logic
// wll be build on top of it
type Service struct {
	Store Store
}

// NewService - which returns a struct of service logic
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - which gets comment based on Id
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retriving Comment Based on Id")
	// ctx = context.WithValue(ctx, "request_id", "unique-string")
	// fmt.Println(ctx.Value("request_id"))
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}
	return cmt, nil

}

func (s *Service) UpdateComment(ctx context.Context, updatedComment Comment, id string) (Comment, error) {
	fmt.Println("Updating Comment")
	cmt, err := s.Store.UpdateComment(ctx, updatedComment, id)
	if err != nil {
		return Comment{}, ErrorUpdatingComment
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorCreatingComment
	}
	return cmt, nil
}
