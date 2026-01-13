package auth

import (
	"context"
	"errors"

	"github.com/priyanshu-samal/miniauth/internal/model"
)

type UserStore interface {
	FindByEmail(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) error
}

type Service struct {
	users UserStore
}

func NewService(users UserStore) *Service {
	return &Service{users: users}
}

func (s *Service) Signup(ctx context.Context, email, password string) error {
	_, err := s.users.FindByEmail(ctx, email)
	if err == nil {
		return errors.New("user already exists")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:    email,
		Password: hash,
	}

	return s.users.Create(ctx, user)
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return GenerateToken(email)
}
