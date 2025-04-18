package user

import (
	"context"
	"go-ent-gin-demo/ent"
)

type Repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) Create(ctx context.Context, input CreateUserInput) (*ent.User, error) {
	return r.client.User.
		Create().
		SetName(input.Name).
		SetAge(input.Age).
		Save(ctx)
}

func (r *Repository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *Repository) DeleteByID(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}
