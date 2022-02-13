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

type IngredientRepositoryInterface interface {
	IsExistIngredient(ingredient string) bool
	InsertIngredient(ingredient string) error
	DeleteIngredient(ingredient string) error
	GetIngredientByName(name string) (models.Ingredients, error)
	GetIngredientIDByName(name string) (int, error)
	GetIngredientByID(id int) (models.Ingredients, error)
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
type ProductRepositoryInterface interface {
	GetProductByID(id int) (models.Product, error)
	GetListOfProdInCategory(catID int) []models.Product
	GetListOfProdBySupplier(suppID int) []models.Product
	GetAllProducts() []models.Product
	InsertToProducts(mp models.Position, productCategoryID int) (int, error)
	UpdateProductById(mp models.Product) error
}
type ProductsCategoriesRepositoryInterface interface {
	CreateCategory(category models.ProductsCategories) (int, error)
	DeleteCategory(id int) error
	GetCategoryByID(id int) (*models.ProductsCategories, error)
	GetAllCategories() ([]models.ProductsCategories, error)
}
type ProductsIngredientsRepositoryInterface interface {
	InsertProductIngredient(productID, ingredientID int) error
	GetIngredientsByProductID(id int) ([]models.ProductsIngredients, error)
	DeleteIngredientByProduct(productID int) error
}
type ProductsSuppliersRepositoryInterface interface {
	IsExistProductSupplier(ps models.ProductsSuppliers) bool
	InsertProductSupplier(ps models.ProductsSuppliers) error
	UpdatePriceByExternalData(price float64, extProdID, extSuppID int) error
	DeleteProductBySupplier(ps models.ProductsSuppliers) error
	GetAllExternalProductIDByExternalSupplierID(extSuppID int) ([]int, error)
	IsExist(prodID, suppID int) bool
}
type SupplierRepositoryInterface interface {
	CreateSupplier(supp models.SupplierForParse, categorySupplierID int) (int, error)
	GetSupplierByID(id int) (models.Supplier, error)
	GetSupplierByName(name string) ([]models.Supplier, error)
	GetAllSuppliers() ([]models.Supplier, error)
	UpdateWorkingHoursByID(ms models.Supplier) error
}
type SuppliersCategoriesRepositoryInterface interface {
	CreateCategory(categories models.SuppliersCategories) error
	GetSupplierCategoryID(name string) (int, error)
}
type UserContactRepositoryInterface interface {
	CreateUserInfo(data models.UserContactData) error
	GetUserInfoByUserID(userID int) ([]models.UserContactData, error)
	GetUserAddressByUSerID(userID int) (string, error)
	UpdateAddress(userID int, newAddress string) error
}
