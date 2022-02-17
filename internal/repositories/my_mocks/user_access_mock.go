package my_mocks

import (
	config "delivery/configs"
	"delivery/internal/auth/services"
	"delivery/internal/models"
	"log"
)

type UserAccessRepositoryMock struct {
	users []*models.UserAccessToken
}

func (u UserAccessRepositoryMock) InsertAccessToken(userToken models.UserAccessToken) error {
	var err error
	var user = models.UserAccessToken{
		UserID:      userToken.UserID,
		AccessToken: userToken.AccessToken,
		CreatedAt:   nil,
		ExpiredAt:   nil,
		Expired:     "false",
	}
	u.users = append(u.users, &user)
	return err
}

func (u UserAccessRepositoryMock) IsExistAccess(userID int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserAccessRepositoryMock) GetAccessTokenByUserID(userID int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserAccessRepositoryMock) GetByAccessToken(accessToken string) (models.UserAccessToken, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserAccessRepositoryMock) UpdateOldAndInsertNewAccessToken(userID int, response models.UserAccessToken) error {
	//TODO implement me
	panic("implement me")
}

func (u UserAccessRepositoryMock) ExpiredAccessToken(userID int) error {
	var err error
	u.users[userID].Expired = "true"
	return err
}

func (u UserAccessRepositoryMock) DeleteNaturallyExpiredAccessToken() {
	//TODO implement me
	panic("implement me")
}

func NewUserAccessRepoMock(cfg *config.Config) *UserAccessRepositoryMock {
	a1, err := services.GenerateToken(1, cfg.AccessLifetimeMinutes, cfg.AccessSecret)
	if err != nil {
		log.Println(err)
	}
	a2, err := services.GenerateToken(2, cfg.AccessLifetimeMinutes, cfg.AccessSecret)

	if err != nil {
		log.Println(err)
	}
	return &UserAccessRepositoryMock{users: []*models.UserAccessToken{
		&models.UserAccessToken{
			UserID:      1,
			AccessToken: a1,
			CreatedAt:   nil,
			ExpiredAt:   nil,
			Expired:     "false",
		},
		&models.UserAccessToken{
			UserID:      2,
			AccessToken: a2,
			CreatedAt:   nil,
			ExpiredAt:   nil,
			Expired:     "false",
		},
	},
	}
}
