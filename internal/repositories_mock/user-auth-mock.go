package repositories_mock

import (
	"delivery/internal/models"
)

type UserRepositoryMock struct {
	users []*models.User
}

//del
//
//func NewUserRepositoryMock() *UserRepositoryMock {
//	p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
//	p2, _ := bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)
//
//	users := []*models.User{
//		&models.User{
//			ID:           1,
//			Email:        "alex-test@example.com",
//			Login:        "Alex",
//			PasswordHash: string(p1),
//		},
//		&models.User{
//			ID:           2,
//			Email:        "mary@example.com",
//			Login:        "Mary",
//			PasswordHash: string(p2),
//		},
//	}
//
//	return &UserRepositoryMock{users: users}
//}
//
//func (r *UserRepositoryMock) GetUserByEmail(email string) (*models.User, error) {
//	for _, user := range r.users {
//		if user.Email == email {
//			return user, nil
//		}
//	}
//
//	return nil, errors.New("user not found")
//}
//
//func (r *UserRepositoryMock) GetUserByID(id int) (*models.User, error) {
//	for _, user := range r.users {
//		if user.ID == id {
//			return user, nil
//		}
//	}
//
//	return nil, errors.New("user not found")
//}
