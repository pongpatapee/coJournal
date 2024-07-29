package service

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"context"

	"github.com/google/uuid"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Create(ctx context.Context, user *entities.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userService) FindAll(ctx context.Context) ([]*entities.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *userService) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *userService) Update(ctx context.Context, user *entities.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}
