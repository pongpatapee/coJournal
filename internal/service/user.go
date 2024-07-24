package service

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"

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

func (s *userService) Create(user *entities.User) error {
	user.ID = uuid.New()
	return s.userRepo.Create(user)
}

func (s *userService) FindAll() ([]*entities.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) FindByID(id uuid.UUID) (*entities.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) Update(user *entities.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) Delete(id uuid.UUID) error {
	return s.userRepo.Delete(id)
}
