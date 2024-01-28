package users

import (
	"context"
	"fmt"
	"github.com/samber/do"
	"time"
)

func ProvideUsersRepository(i *do.Injector) (*Repository, error) {
	return NewUsersRepository(i), nil
}

func NewUsersRepository(_ *do.Injector) *Repository {
	return &Repository{}
}

type Repository struct {
}

func (r *Repository) GetUser(ctx context.Context, id int) (*User, error) {
	timeToGet := 3 * time.Second
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf(`can't get user with id %d`, id)
	case <-time.After(timeToGet):
		return &User{}, nil
	}
}
