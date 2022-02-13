package repository_innterfaces

import "delivery/internal/models"

type IngredientRepositoryInterface interface {
	IsExistIngredient(ingredient string) bool
	InsertIngredient(ingredient string) error
	DeleteIngredient(ingredient string) error
	GetIngredientByName(name string) (models.Ingredients, error)
	GetIngredientIDByName(name string) (int, error)
	GetIngredientByID(id int) (models.Ingredients, error)
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
