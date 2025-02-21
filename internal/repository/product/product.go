package repository

import (
	"database/sql"
	"log"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProductRepository(db *sql.DB) *ProductRepository {
	defaultDb := &ProductRepository{
		db: db,
	}
	return defaultDb
}

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) GetProduct() (productMap map[int]model.Product, err error) {
	rows, err := r.db.Query("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products")
	if err != nil {
		return nil, errorHandler.GetErrDatabase(errorHandler.PRODUCT)
	}
	productMap = make(map[int]model.Product)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.ProductAttributes.ProductCode, &product.ProductAttributes.Description, &product.ProductAttributes.Width, &product.ProductAttributes.Height, &product.ProductAttributes.Length, &product.ProductAttributes.NetWeight, &product.ProductAttributes.ExpirationRate, &product.ProductAttributes.RecommendedFreezingTemperature, &product.ProductAttributes.FreezingRate, &product.ProductAttributes.ProductTypeID, &product.ProductAttributes.SellerID)
		if err != nil {
			return nil, errorHandler.GetErrParsingData(errorHandler.PRODUCT)
		}
		productMap[product.ID] = product
	}
	return
}

func (r *ProductRepository) GetProductByID(id int) (product model.Product, err error) {
	row := r.db.QueryRow("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products WHERE id = ?", id)
	err = row.Scan(&product.ID, &product.ProductAttributes.ProductCode, &product.ProductAttributes.Description, &product.ProductAttributes.Width, &product.ProductAttributes.Height, &product.ProductAttributes.Length, &product.ProductAttributes.NetWeight, &product.ProductAttributes.ExpirationRate, &product.ProductAttributes.RecommendedFreezingTemperature, &product.ProductAttributes.FreezingRate, &product.ProductAttributes.ProductTypeID, &product.ProductAttributes.SellerID)
	if err != nil {
		return model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	return product, nil
}

func (r *ProductRepository) GetProductRecord() (map[int]model.ProductRecordCount, error) {
	rows, err := r.db.Query("SELECT prod.id as product_id, prod.description, COUNT(*) as records_count FROM products as prod INNER JOIN product_records as pr ON prod.id = pr.product_id GROUP BY prod.id")
	if err != nil {
		return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD)
	}
	defer rows.Close()

	var productRecords = make(map[int]model.ProductRecordCount)
	for rows.Next() {
		var productRecord model.ProductRecordCount
		err := rows.Scan(&productRecord.ProductID, &productRecord.Description, &productRecord.Count)

		if err != nil {
			return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD)
		}
		productRecords[productRecord.ProductID] = productRecord
	}

	return productRecords, nil
}

func (r *ProductRepository) GetProductRecordByID(id int) (model.ProductRecordCount, error) {
	row := r.db.QueryRow("SELECT prod.id as product_id, prod.description, COUNT(*) as records_count FROM products as prod INNER JOIN product_records as pr ON prod.id = pr.product_id GROUP BY prod.id HAVING prod.id = ?", id)
	var productRecordCount model.ProductRecordCount
	err := row.Scan(&productRecordCount.ProductID, &productRecordCount.Description, &productRecordCount.Count)
	if err != nil {
		return model.ProductRecordCount{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD)
	}
	return productRecordCount, nil
}

func (r *ProductRepository) productCodeExists(productCode string) bool {
	row := r.db.QueryRow("SELECT COUNT(*) FROM products WHERE product_code = ?", productCode)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (r *ProductRepository) CreateProduct(productAtrributes *model.ProductAttributes) (prod model.Product, err error) {

	if r.productCodeExists(*productAtrributes.ProductCode) {
		return model.Product{}, errorHandler.GetErrAlreadyExists(errorHandler.PRODUCT)
	}
	row, err := r.db.Exec("INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", productAtrributes.ProductCode, productAtrributes.Description, productAtrributes.Width, productAtrributes.Height, productAtrributes.Length, productAtrributes.NetWeight, productAtrributes.ExpirationRate, productAtrributes.RecommendedFreezingTemperature, productAtrributes.FreezingRate, productAtrributes.ProductTypeID, productAtrributes.SellerID)

	if err != nil {
		return model.Product{}, errorHandler.GetErrParsingData(errorHandler.PRODUCT)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Product{}, errorHandler.GetErrDatabase(errorHandler.PRODUCT)
	}

	var newProduct model.Product
	newProduct.ID = int(id)
	newProduct.ProductAttributes = *productAtrributes

	return newProduct, nil
}

func (r *ProductRepository) DeleteProduct(id int) (err error) {
	exist, err := r.registerExists(id)

	if err != nil {
		return err
	}

	if !exist {
		return errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	_, err = r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(id int, productAtrributesPtr *model.ProductAttributes) (product *model.Product, err error) {

	exist, err := r.registerExists(id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	if productAtrributesPtr == nil {
		return nil, errorHandler.GetErrInvalidData(errorHandler.PRODUCT)
	}

	if productAtrributesPtr.ProductCode != nil && r.productCodeExists(*productAtrributesPtr.ProductCode) {
		return nil, errorHandler.GetErrAlreadyExists(errorHandler.PRODUCT)

	}

	var patchedProduct model.ProductAttributes
	product = &model.Product{}
	err = r.db.QueryRow("SELECT product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id FROM products WHERE id = ?", id).Scan(&patchedProduct.ProductCode, &patchedProduct.Description, &patchedProduct.Width, &patchedProduct.Height, &patchedProduct.Length, &patchedProduct.NetWeight, &patchedProduct.ExpirationRate, &patchedProduct.RecommendedFreezingTemperature, &patchedProduct.ProductTypeID, &patchedProduct.SellerID)

	if err != nil {
		return nil, errorHandler.GetErrParsingData(errorHandler.PRODUCT)
	}

	if productAtrributesPtr.ProductCode != nil {
		patchedProduct.ProductCode = productAtrributesPtr.ProductCode
	}
	if productAtrributesPtr.Description != nil {
		patchedProduct.Description = productAtrributesPtr.Description
	}
	if productAtrributesPtr.Width != nil {
		patchedProduct.Width = productAtrributesPtr.Width
	}
	if productAtrributesPtr.Height != nil {
		patchedProduct.Height = productAtrributesPtr.Height
	}
	if productAtrributesPtr.Length != nil {
		patchedProduct.Length = productAtrributesPtr.Length
	}
	if productAtrributesPtr.NetWeight != nil {
		patchedProduct.NetWeight = productAtrributesPtr.NetWeight
	}
	if productAtrributesPtr.ExpirationRate != nil {
		patchedProduct.ExpirationRate = productAtrributesPtr.ExpirationRate
	}
	if productAtrributesPtr.RecommendedFreezingTemperature != nil {
		patchedProduct.RecommendedFreezingTemperature = productAtrributesPtr.RecommendedFreezingTemperature
	}
	if productAtrributesPtr.FreezingRate != nil {
		patchedProduct.FreezingRate = productAtrributesPtr.FreezingRate
	}
	if productAtrributesPtr.ProductTypeID != nil {
		patchedProduct.ProductTypeID = productAtrributesPtr.ProductTypeID
	}

	// Update the product in the repository after all fields have been patched
	_, err = r.db.Exec("UPDATE products SET product_code = ?, description = ?, width = ?, height = ?, length = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, product_type_id = ?, seller_id = ? WHERE id = ?", patchedProduct.ProductCode, patchedProduct.Description, patchedProduct.Width, patchedProduct.Height, patchedProduct.Length, patchedProduct.NetWeight, patchedProduct.ExpirationRate, patchedProduct.RecommendedFreezingTemperature, patchedProduct.ProductTypeID, patchedProduct.SellerID, id)
	if err != nil {
		log.Println(err)
		return nil, errorHandler.GetErrInvalidData(errorHandler.PRODUCT)
	}
	product.ID = id
	product.ProductAttributes = patchedProduct
	return product, nil
}

func (r *ProductRepository) registerExists(id int) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM products WHERE ID = ?)"
	err := r.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false, errorHandler.GetErrDatabase(errorHandler.PRODUCT)
	}
	return exist, nil
}

/* func (r *ProductRepository) registerExists(id int) (bool, error) {
	var exist int
	query := "SELECT COUNT(*) FROM products WHERE ID = ?"
	err := r.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false, errorHandler.GetErrDatabase(errorHandler.PRODUCT)
	}
	if exist == 0 {
		return false, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	fmt.Println(exist)
	return exist > 0, nil
} */
