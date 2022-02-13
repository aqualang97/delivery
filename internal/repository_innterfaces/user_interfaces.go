package repository_innterfaces

import "delivery/internal/models"

//go:generate mockgen -source=service.go -destination=mocks/user.go

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

type OrderProductRepositoryInterface interface {
	InsertToOrdersProducts(mo models.OrderProducts) (int, error)
	UpdateNumbersByProductAndOrderID(mo models.OrderProducts) error
	GetAllProductsByOrderID(orderID int) ([]models.OrderProducts, error)
	DeleteProduct(order models.OrderProducts) error
	DeleteAll(order models.OrderProducts) error
}
type OrderRepositoryInterface interface {
	InsertToOrders(mo models.Order) (int, error)
	UpdateOrdersByID(mo *models.Order) error
	UpdateOrdersByUserID(mo *models.Order) error
	GetOrderByID(id int) (models.Order, error)
	GetOrderByUserIDNotPaidNotCompleted(userID int) (models.Order, error)
	DeleteOrderByUserID(userID int) error
}

type UserContactRepositoryInterface interface {
	CreateUserInfo(data models.UserContactData) error
	GetUserInfoByUserID(userID int) ([]models.UserContactData, error)
	GetUserAddressByUSerID(userID int) (string, error)
	UpdateAddress(userID int, newAddress string) error
}
