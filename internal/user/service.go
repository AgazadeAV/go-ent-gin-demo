package user

import (
	"context"
	"go-ent-gin-demo/ent"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, input CreateUserInput) (*ent.User, error) {
	return s.repo.Create(ctx, input)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteByID(ctx, id)
}
