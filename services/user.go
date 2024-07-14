package services

import (
	"sync"

	"github.com/google/uuid"
)

type UserID string

type User struct {
	// CreateDate  time.Time `json:"create_date"`
	// UpdatedDate time.Time `json:"updated_date"`
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UserService struct {
	db map[UserID]User
	sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		db: make(map[UserID]User),
	}
}

func (us *UserService) Create(user User) User {
	us.Lock()
	defer us.Unlock()

	id := UserID(uuid.New().String())
	user.ID = string(id)
	us.db[id] = user
	return user
}

func (us *UserService) GetAll() []User {
	us.Lock()
	defer us.Unlock()

	users := make([]User, 0, len(us.db))
	for _, user := range us.db {
		users = append(users, user)
	}
	return users
}

func (us *UserService) Get(id UserID) (User, bool) {
	us.Lock()
	defer us.Unlock()

	user, ok := us.db[id]

	return user, ok
}

func (us *UserService) Update(id UserID, newUser User) (User, bool) {
	us.Lock()
	defer us.Unlock()

	_, ok := us.db[id]

	if !ok {
		return User{}, false
	}

	newUser.ID = string(id)
	us.db[id] = newUser

	return newUser, true
}

func (us *UserService) Delete(id UserID) bool {
	us.Lock()
	defer us.Unlock()

	_, ok := us.db[id]
	if !ok {
		return false
	}

	delete(us.db, id)
	return true
}
