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

func (p ProductDBRepository) GetProductByID(id int) (models.Position, error) {
	var product models.Position
	var ingredient string
	var listIngr []string

	err := p.conn.QueryRow(
		"SELECT p.id, p.name, s.id, s.name, p.category, pc.id, p.external_id, pc.name, ps.price, ps.image FROM products AS p LEFT JOIN products_categories pc on pc.id = p.category left join products_suppliers ps on p.id = ps.product_id left join suppliers s on ps.supplier_id = s.id  WHERE p.id=? and p.external_id=ps.external_product_id",
		id).Scan(&product.ID, &product.Name, &product.SupplierId, &product.SupplierName, &product.CategoryNum, &product.CategoryNum, &product.ExternalID, &product.Type, &product.Price, &product.Image)
	if err != nil {
		log.Println(err)
		return product, err
	}

	rowsI, err := p.conn.Query(
		"SELECT i.name FROM ingredients as i LEFT JOIN products_ingredients as pi on i.id = pi.ingredient_id  WHERE pi.product_id=?", product.ID)
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rowsI.Close()
	for rowsI.Next() {
		err := rowsI.Scan(&ingredient)
		if err != nil {
			log.Println(err)
			return product, err
		}
		listIngr = append(listIngr, ingredient)
	}

	product.Ingredients = listIngr
	return product, err
}
func (p ProductDBRepository) GetListOfProdInCategory(catID int) []models.Position {
	var product models.Position
	var listProd []models.Position
	var ingredientName models.ProductsIngredientsName
	var ingredientNameList []models.ProductsIngredientsName
	rowsI, err := p.conn.Query(
		"SELECT pi.product_id, i.name FROM products_ingredients as pi LEFT JOIN ingredients i on i.id = pi.ingredient_id ORDER BY product_id")
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rowsI.Close()

	rows, err := p.conn.Query(
		"SELECT products.id, products.name, pc.name, pc.id, products.external_id, s.name, ps.price, ps.image, ps.supplier_id, ps.external_supplier_id FROM products LEFT JOIN products_categories pc ON pc.id = products.category LEFT JOIN products_suppliers ps ON products.external_id = ps.external_product_id LEFT JOIN suppliers s ON ps.external_supplier_id = s.external_id WHERE pc.id = ?", catID)
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.CategoryNum, &product.ExternalID, &product.SupplierName, &product.Price, &product.Image, &product.SupplierId, &product.ExternalSuppId)
		if err != nil {

			return listProd
		}
		listProd = append(listProd, product)

	}
	for rowsI.Next() {
		err := rowsI.Scan(&ingredientName.ProductId, &ingredientName.Ingredient)
		if err != nil {
			log.Println(err)
			return listProd
		}
		ingredientNameList = append(ingredientNameList, ingredientName)
	}
	for i, _ := range ingredientNameList {
		for j, _ := range listProd {
			if ingredientNameList[i].ProductId == listProd[j].ID {
				//log.Println(ingredientNameList[i].ProductId, listProd[j].ID)

				listProd[j].Ingredients = append(listProd[j].Ingredients, ingredientNameList[i].Ingredient)
			} else {
				continue
			}
		}

	}
	return listProd
}
func (p ProductDBRepository) GetListOfProdBySupplier(suppID int) []models.Position {
	var product models.Position
	var listProd []models.Position
	var ingredientName models.ProductsIngredientsName
	var ingredientNameList []models.ProductsIngredientsName
	rowsI, err := p.conn.Query(
		"SELECT pi.product_id, i.name FROM products_ingredients as pi LEFT JOIN ingredients i on i.id = pi.ingredient_id ORDER BY product_id")
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rowsI.Close()

	rows, err := p.conn.Query(
		"SELECT products.id, products.name, pc.name, pc.id, products.external_id, s.name, ps.price, ps.image, ps.supplier_id, ps.external_supplier_id FROM products LEFT JOIN products_categories pc ON pc.id = products.category LEFT JOIN products_suppliers ps ON products.external_id = ps.external_product_id LEFT JOIN suppliers s ON ps.external_supplier_id = s.external_id WHERE ps.supplier_id=?", suppID)
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.CategoryNum, &product.ExternalID, &product.SupplierName, &product.Price, &product.Image, &product.SupplierId, &product.ExternalSuppId)
		if err != nil {
			log.Println(err)
			return listProd
		}
		//fmt.Println(product)
		listProd = append(listProd, product)
	}
	for rowsI.Next() {
		err := rowsI.Scan(&ingredientName.ProductId, &ingredientName.Ingredient)
		if err != nil {
			log.Println(err)
			return listProd
		}
		ingredientNameList = append(ingredientNameList, ingredientName)
	}
	for i, _ := range ingredientNameList {
		for j, _ := range listProd {
			if ingredientNameList[i].ProductId == listProd[j].ID {
				//log.Println(ingredientNameList[i].ProductId, listProd[j].ID)

				listProd[j].Ingredients = append(listProd[j].Ingredients, ingredientNameList[i].Ingredient)
			} else {
				continue
			}
		}

	}

	return listProd
}
func (p ProductDBRepository) GetAllProducts() []models.Position {
	var product models.Position
	var listProd []models.Position
	var piProductID int
	var ingredient string
	var listIngr []string
	rowsI, err := p.conn.Query(
		"SELECT pi.product_id, i.name FROM products_ingredients as pi LEFT JOIN ingredients i on i.id = pi.ingredient_id ORDER BY product_id")
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rowsI.Close()

	rows, err := p.conn.Query(
		"SELECT products.id, products.name, pc.name, pc.id, products.external_id, ps.price, ps.image, ps.supplier_id, s.name, ps.external_supplier_id AS prod_image FROM products LEFT JOIN products_categories pc ON pc.id = products.category LEFT JOIN products_suppliers ps ON products.external_id = ps.external_product_id LEFT JOIN suppliers s ON ps.external_supplier_id = s.external_id")
	if err != nil {
		p.logger.Error("GetListOfProdBySupplier \n", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.CategoryNum, &product.ExternalID, &product.Price, &product.Image, &product.SupplierId, &product.SupplierName, &product.ExternalSuppId)
		if err != nil {
			log.Println(err)
			return listProd
		}
		//fmt.Println(product)
		for rowsI.Next() {
			err := rowsI.Scan(&piProductID, &ingredient)
			if err != nil {
				log.Println(err)
				return listProd
			}
			if piProductID != product.ID {
				break
			}
			listIngr = append(listIngr, ingredient)
		}
		product.Ingredients = listIngr
		listProd = append(listProd, product)
		listIngr = nil
	}

	return listProd
}
func (p ProductDBRepository) InsertToProducts(mp models.Position, productCategoryID int) (int, error) {
	var exist bool
	var id int64
	err := p.conn.QueryRow("SELECT EXISTS(SELECT * FROM products WHERE name=? AND external_id=?)",
		mp.Name, mp.ExternalID).Scan(&exist)

	if err != nil {
		log.Println(err)
		return int(id), err
	}

	if !exist {
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
