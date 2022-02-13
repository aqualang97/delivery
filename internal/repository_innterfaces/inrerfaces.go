package repository_innterfaces

import "delivery/internal/models"

type UserRepositoryInterface interface {
	GetUserById(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
	CreateUser(user *models.RegistrationRequest, passwordHash string) (int, error)
	UpdateUserById(user *models.User) error
}

type UserAccessTokenRepositoryInterface interface {
	InsertAccessToken(userToken models.UserAccessToken) error
	IsExistAccess(userID int) (bool, error)
	GetAccessTokenByUserID(userID int) (string, error)
	GetByAccessToken(accessToken string) (models.UserAccessToken, error)
	UpdateOldAndInsertNewAccessToken(userID int, response models.UserAccessToken) error
	ExpiredAccessToken(userID int) error
	DeleteNaturallyExpiredAccessToken()
}

type UserRefreshTokenRepositoryInterface interface {
	InsertRefreshToken(userToken models.UserRefreshToken) error
	IsExistRefresh(userID int) (bool, error)
	GetRefreshTokenByUserID(userID int) (string, error)
	GetByRefreshToken(refreshToken string) (models.UserRefreshToken, error)
	UpdateOldAndInsertNewRefreshToken(userID int, response models.UserRefreshToken) error
	ExpiredRefreshToken(userID int) error
	DeleteNaturallyExpiredRefreshToken()
}

func n() {

}