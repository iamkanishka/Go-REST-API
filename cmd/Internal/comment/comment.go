package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("Failed to Fetch comment by Id")
	ErrorNotImplemented  = errors.New("Method not Implemented")
)

// Comment - Representation Structure of Comment fo
// Service
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

// Store -  this interface defines all the
// methods that out service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
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
	fmt.Println("Retriving Cooment Based on Id")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}
	return cmt, nil

}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}
