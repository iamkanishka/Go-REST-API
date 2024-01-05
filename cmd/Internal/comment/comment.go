package comment

import (
	"context"
	"fmt"
)

// Comment - Representation Structure of Comment fo
// Service
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

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
		return Comment{}, nil
	}
	return cmt, nil

}
