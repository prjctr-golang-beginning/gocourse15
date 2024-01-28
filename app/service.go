package app

import (
	"context"
	"github.com/samber/do"
	"gocourse15/pkg/users"
	"time"
)

type Service struct {
	usersRepository *users.Repository
}

func (s *Service) Do(ctx context.Context) error {
	userId := 0
	tCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	_, err := s.usersRepository.GetUser(tCtx, userId)

	return err
}

func ProvideService(i *do.Injector) (*Service, error) {
	return &Service{
		usersRepository: do.MustInvoke[*users.Repository](i),
	}, nil
}
