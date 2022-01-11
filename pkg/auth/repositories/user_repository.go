package repositories

import (
	"delivery/internal/models"
	db "delivery/internal/repositories/database"
	authModels "delivery/pkg/auth/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	users []*authModels.User
}

func NewUserRepository() *UserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)
	users := []*authModels.User{
		&authModels.User{
			ID:       1,
			Email:    "Alex@example.com",
			Name:     "Alex",
			Password: string(p1),
		},
		&authModels.User{
			ID:       2,
			Email:    "mary@example.com",
			Name:     "Mary",
			Password: string(p2),
		},
	}
	return &UserRepository{users: users}
}
func NewUserRepositoryReg(email, name, password string) *UserRepository {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := []*authModels.User{
		{
			Email:    email,
			Name:     name,
			Password: string(passwordHash),
		},
	}

	return &UserRepository{users: user}
}

func NewUserRepositoryLogin(email, password string) *UserRepository {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := []*authModels.User{
		{
			Email:    email,
			Password: string(passwordHash),
		},
	}
	return &UserRepository{users: user}
}
func (r *UserRepository) GetUserByEmail(email string) (*authModels.User, error) {

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *UserRepository) GetUserByEmailFromDB(email string) (models.User, error) {
	user, err := db.UserDBRepository{}.GetUserByEmail(email)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil

}
func (r *UserRepository) GetUserByID(id int) (*authModels.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
