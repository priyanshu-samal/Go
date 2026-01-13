package auth

import (
	"context"
	"errors"

	"github.com/priyanshu-samal/miniauth/internal/model"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:    email,
		Password: string(hash),
	}

	return s.users.Create(ctx, user)
}

func (s *Service) Login(ctx context.Context, email, password string) error {
	user, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
