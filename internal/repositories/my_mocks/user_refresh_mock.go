package my_mocks

import (
	config "delivery/configs"
	"delivery/internal/auth/services"
	"delivery/internal/models"
	"log"
)

type UserRefreshRepositoryMock struct {
	users []*models.UserRefreshToken
}

func (u UserRefreshRepositoryMock) InsertRefreshToken(userToken models.UserRefreshToken) error {
	var err error
	var user = models.UserRefreshToken{
		UserID:       userToken.UserID,
		RefreshToken: userToken.RefreshToken,
		CreatedAt:    nil,
		ExpiredAt:    nil,
		Expired:      "false",
	}
	u.users = append(u.users, &user)
	return err
}

func (u UserRefreshRepositoryMock) IsExistRefresh(userID int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRefreshRepositoryMock) GetRefreshTokenByUserID(userID int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRefreshRepositoryMock) GetByRefreshToken(refreshToken string) (models.UserRefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRefreshRepositoryMock) UpdateOldAndInsertNewRefreshToken(userID int, response models.UserRefreshToken) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRefreshRepositoryMock) ExpiredRefreshToken(userID int) error {
	var err error
	u.users[userID].Expired = "true"
	return err
}

func (u UserRefreshRepositoryMock) DeleteNaturallyExpiredRefreshToken() {
	//TODO implement me
	panic("implement me")
}

func NewUserRefreshRepoMock(cfg *config.Config) *UserRefreshRepositoryMock {
	r1, err := services.GenerateToken(1, cfg.RefreshLifetimeMinutes, cfg.RefreshSecret)
	if err != nil {
		log.Println(err)
	}

	r2, err := services.GenerateToken(2, cfg.RefreshLifetimeMinutes, cfg.RefreshSecret)

	if err != nil {
		log.Println(err)
	}
	return &UserRefreshRepositoryMock{users: []*models.UserRefreshToken{
		&models.UserRefreshToken{
			UserID:       1,
			RefreshToken: r1,
			CreatedAt:    nil,
			ExpiredAt:    nil,
			Expired:      "false",
		},
		&models.UserRefreshToken{
			UserID:       2,
			RefreshToken: r2,
			CreatedAt:    nil,
			ExpiredAt:    nil,
			Expired:      "false",
		},
	},
	}
}
