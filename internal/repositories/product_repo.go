package repositories

import (
	"database/sql"
	"delivery/internal/models"
	"github.com/aqualang97/logger/v4"
	"log"
)

type ProductDBRepository struct {
	conn   *sql.DB
	TX     *sql.Tx
	logger *logger.Logger
}

func (p ProductDBRepository) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	err := p.conn.QueryRow(
		"SELECT id, name, category, external_id FROM products WHERE id = ?",
		id).Scan(&product.ID, &product.Name, &product.Category, &product.ExternalID)
	if err != nil {
		log.Println(err)
	}
	return product, err
}
func (p ProductDBRepository) GetListOfProdInCategory(catID int) []models.Product {
	var product models.Product
	var listProd []models.Product

	//rows, err := p.conn.Query("SELECT products.id, products.name, products.external_id FROM products INNER JOIN products_categories on products.category =?", catID)
	rows, err := p.conn.Query("SELECT products.id, products.name, products.external_id FROM products WHERE products.category =?", catID)

	if err != nil {
		p.logger.Error("GetListOfProdInCategory \n", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.ExternalID)
		if err != nil {
			log.Println(err)
			return listProd
		}
		listProd = append(listProd, product)
	}

	return listProd
}
func (p ProductDBRepository) GetListOfProdBySupplier(suppID int) []models.Product {
	var product models.Product
	var listProd []models.Product

	rows, err := p.conn.Query("SELECT products_suppliers.product_id FROM products_suppliers where supplier_id=?", suppID)
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&product.ID)
		if err != nil {
			log.Println(err)
			return listProd

		}
		err = p.conn.QueryRow(
			"SELECT name, category, external_id FROM products WHERE id = ?",
			product.ID).Scan(&product.Name, &product.Category, &product.ExternalID)
		if err != nil {
			log.Println(err)
		}

		listProd = append(listProd, product)
	}

	return listProd
}
func (p ProductDBRepository) GetAllProducts() []models.Position {
	var product models.Position
	var listProd []models.Position
	//	var ingridient string
	//	var litIngr []string
	rows, err := p.conn.Query(
		"SELECT products.id, products.name, pc.name, products.external_id, ps.price, ps.image AS prod_image FROM products LEFT JOIN products_categories pc ON pc.id = products.category LEFT JOIN products_suppliers ps ON products.external_id = ps.external_product_id LEFT JOIN suppliers s ON products.external_id = s.external_id")
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.ExternalID, &product.Price, &product.Image)
		if err != nil {
			log.Println(err)
			return listProd

		}
		rowsI, err := p.conn.Query(
			"SELECT products.id, products.name, pc.name, products.external_id, ps.price, ps.image AS prod_image FROM products LEFT JOIN products_categories pc ON pc.id = products.category LEFT JOIN products_suppliers ps ON products.external_id = ps.external_product_id LEFT JOIN suppliers s ON products.external_id = s.external_id")
		defer rowsI.Close()

		listProd = append(listProd, product)
	}
	return listProd
}
func (p ProductDBRepository) InsertToProducts(mp models.Position, productCategoryID int) (int, error) {
	res, err := p.conn.Exec(
		"INSERT products(name, category, external_id) VALUES(?, ?, ?)",
		mp.Name, productCategoryID, mp.ExternalID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(id), err
}

func (p ProductDBRepository) UpdateProductById(mp models.Product) error {
	var err error // ?
	// Пока не представляю что можно тут обновлять
	//
	return err
}

func (p ProductDBRepository) DeleteProductByExternalID(name string, externalID int) error {
	_, err := p.conn.Exec("DELETE FROM products WHERE name=? AND external_id=?", name, externalID)
	if err != nil {
		log.Println(err)
	}
	return err
}
func NewProductRepo(conn *sql.DB, TX *sql.Tx, logger *logger.Logger) *ProductDBRepository {
	return &ProductDBRepository{conn: conn, TX: TX, logger: logger}
}
