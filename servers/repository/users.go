package repository

import (
	"context"

	"servers/domain"
	"servers/pkg/rds"
)

type Users struct {
	r rds.RDS
}

func NewUsers(r rds.RDS) *Users {
	return &Users{r: r}
}

func (u *Users) Store(ctx context.Context, user *domain.User) (int64, error) {
	q := `INSERT INTO users (name, description, updated_at, created_at)
VALUES (?, ?, NOW(), NOW())`
	result, err := u.r.ExecContext(ctx, q, user.Name, user.Description)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (u *Users) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	var user *domain.User
	if err := u.r.SelectContext(ctx, user, "SELECT * FROM users WHERE id = ?", id); err != nil {
		return user, err
	}
	return user, nil
}
