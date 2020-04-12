package usecase

import (
	"context"

	"servers/domain"
)

type usersRepository interface {
	Store(ctx context.Context, user *domain.User) (int64, error)
	FindByID(ctx context.Context, id int64) (*domain.User, error)
}

type Users struct {
	repository usersRepository
}

func NewUsers(r usersRepository) *Users {
	return &Users{repository: r}
}

func (u *Users) Add(ctx context.Context, user *domain.User) (int64, error) {
	return u.repository.Store(ctx, user)
}

func (u *Users) Get(ctx context.Context, id int64) (*domain.User, error) {
	return u.repository.FindByID(ctx, id)
}
