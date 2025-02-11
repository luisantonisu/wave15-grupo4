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

func (productRepository *ProductRepository) GetProduct() (productMap map[int]model.Product, err error) {
	rows, err := productRepository.db.Query("SELECT id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id FROM product")
	if err != nil {
		return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	productMap = make(map[int]model.Product)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.ProductAtrributes.ProductCode, &product.ProductAtrributes.Description, &product.ProductAtrributes.Width, &product.ProductAtrributes.Height, &product.ProductAtrributes.NetWeight, &product.ProductAtrributes.ExpirationRate, &product.ProductAtrributes.RecommendedFreezingTemperature, &product.ProductAtrributes.ProductTypeID, &product.ProductAtrributes.SellerID)
		if err != nil {
			return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
		}
		productMap[product.ID] = product
	}
	return
}

func (productRepository *ProductRepository) GetProductByID(id int) (product model.Product, err error) {
	row := productRepository.db.QueryRow("SELECT id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id FROM product WHERE id = ?", id)
	err = row.Scan(&product.ID, &product.ProductAtrributes.ProductCode, &product.ProductAtrributes.Description, &product.ProductAtrributes.Width, &product.ProductAtrributes.Height, &product.ProductAtrributes.NetWeight, &product.ProductAtrributes.ExpirationRate, &product.ProductAtrributes.RecommendedFreezingTemperature, &product.ProductAtrributes.ProductTypeID, &product.ProductAtrributes.SellerID)
	if err != nil {
		return model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	return product, nil
}

func (productRepository *ProductRepository) GetProductRecord() (map[int]model.ProductRecordCount, error) {
	rows, err := productRepository.db.Query("SELECT prod.id as product_id, prod.description, COUNT(*) as records_count FROM product as prod INNER JOIN product_records as pr ON prod.id = pr.product_id GROUP BY prod.id")
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

func (productRepository *ProductRepository) GetProductRecordByID(id int) (model.ProductRecordCount, error) {
	row := productRepository.db.QueryRow("SELECT prod.id as product_id, prod.description, COUNT(*) as records_count FROM product as prod INNER JOIN product_records as pr ON prod.id = pr.product_id GROUP BY prod.id HAVING prod.id = ?", id)
	var productRecordCount model.ProductRecordCount
	err := row.Scan(&productRecordCount.ProductID, &productRecordCount.Description, &productRecordCount.Count)
	if err != nil {
		return model.ProductRecordCount{}, err
	}
	return productRecordCount, nil
}

func (productRepository *ProductRepository) productCodeExists(productCode string) bool {
	row := productRepository.db.QueryRow("SELECT COUNT(*) FROM product WHERE product_code = ?", productCode)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (productRepository *ProductRepository) CreateProduct(productAtrributes *model.ProductAtrributes) (err error) {

	if productRepository.productCodeExists(productAtrributes.ProductCode) {
		return errorHandler.GetErrAlreadyExists(errorHandler.PRODUCT)
	}
	_, err = productRepository.db.Exec("INSERT INTO product (product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", productAtrributes.ProductCode, productAtrributes.Description, productAtrributes.Width, productAtrributes.Height, productAtrributes.NetWeight, productAtrributes.ExpirationRate, productAtrributes.RecommendedFreezingTemperature, productAtrributes.ProductTypeID, productAtrributes.SellerID)

	if err != nil {
		return err
	}

	return nil
}

func (productRepository *ProductRepository) DeleteProduct(id int) (err error) {
	_, err = productRepository.registerExists(id)
	if err != nil {
		return err
	}
	_, err = productRepository.db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		return errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	return nil
}

func (productRepository *ProductRepository) UpdateProduct(id int, productAtrributesPtr *model.ProductAtrributesPtr) (product *model.Product, err error) {

	_, err = productRepository.registerExists(id)
	if err != nil {
		return nil, err
	}

	if productAtrributesPtr == nil {
		return nil, errorHandler.GetErrInvalidData(errorHandler.PRODUCT)
	}

	if productRepository.productCodeExists(*productAtrributesPtr.ProductCode) {
		return nil, errorHandler.GetErrAlreadyExists(errorHandler.PRODUCT)
	}
	var patchedProduct model.ProductAtrributes
	product = &model.Product{}
	err = productRepository.db.QueryRow("SELECT product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id FROM product WHERE id = ?", id).Scan(&patchedProduct.ProductCode, &patchedProduct.Description, &patchedProduct.Width, &patchedProduct.Height, &patchedProduct.NetWeight, &patchedProduct.ExpirationRate, &patchedProduct.RecommendedFreezingTemperature, &patchedProduct.ProductTypeID, &patchedProduct.SellerID)

	if err != nil {
		return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}

	if productAtrributesPtr.ProductCode != nil {
		patchedProduct.ProductCode = *productAtrributesPtr.ProductCode
	}
	if productAtrributesPtr.Description != nil {
		patchedProduct.Description = *productAtrributesPtr.Description
	}
	if productAtrributesPtr.Width != nil {
		patchedProduct.Width = *productAtrributesPtr.Width
	}
	if productAtrributesPtr.Height != nil {
		patchedProduct.Height = *productAtrributesPtr.Height
	}
	if productAtrributesPtr.Length != nil {
		patchedProduct.Length = *productAtrributesPtr.Length
	}
	if productAtrributesPtr.NetWeight != nil {
		patchedProduct.NetWeight = *productAtrributesPtr.NetWeight
	}
	if productAtrributesPtr.ExpirationRate != nil {
		patchedProduct.ExpirationRate = *productAtrributesPtr.ExpirationRate
	}
	if productAtrributesPtr.RecommendedFreezingTemperature != nil {
		patchedProduct.RecommendedFreezingTemperature = *productAtrributesPtr.RecommendedFreezingTemperature
	}
	if productAtrributesPtr.FreezingRate != nil {
		patchedProduct.FreezingRate = *productAtrributesPtr.FreezingRate
	}
	if productAtrributesPtr.ProductTypeID != nil {
		patchedProduct.ProductTypeID = *productAtrributesPtr.ProductTypeID
	}

	// Update the product in the repository after all fields have been patched
	_, err = productRepository.db.Exec("UPDATE product SET product_code = ?, description = ?, width = ?, height = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, product_type_id = ?, seller_id = ? WHERE id = ?", patchedProduct.ProductCode, patchedProduct.Description, patchedProduct.Width, patchedProduct.Height, patchedProduct.NetWeight, patchedProduct.ExpirationRate, patchedProduct.RecommendedFreezingTemperature, patchedProduct.ProductTypeID, patchedProduct.SellerID, id)
	if err != nil {
		log.Println(err)
		return nil, errorHandler.GetErrInvalidData(errorHandler.PRODUCT)
	}
	product.ID = id
	product.ProductAtrributes = patchedProduct
	return product, nil
}

func (productRepository *ProductRepository) registerExists(id int) (bool, error) {
	var exist bool
	query := "SELECT EXISTS(SELECT 1 FROM product WHERE ID = ?)"
	err := productRepository.db.QueryRow(query, id).Scan(&exist)
	if err != nil {
		return false, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
	}
	return exist, nil
}
