package repositories

import (
	"strings"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func ProductDBInstance(db *gorm.DB) domain.IProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (repo ProductRepo) CreateProducts(products []models.Product) error {
	if err := repo.db.Create(products).Error; err != nil {
		return err
	}

	return nil
}

func (repo ProductRepo) GetProducts(getCategoriesParams types.GetCategoriesParams) []models.Product {
	var Products []models.Product

	query := `select products.id, products.name, products.description, products.category_id, products.sub_category_id, products.icon, products.price, products.quantity, products.discount, products.status from products`

	var params []interface{}

	if getCategoriesParams.CategoryID != 0 && getCategoriesParams.SubCategoryID == 0 {
		query += " WHERE products.category_id =?"
		params = append(params, getCategoriesParams.CategoryID)
	} else if getCategoriesParams.CategoryID != 0 && getCategoriesParams.SubCategoryID != 0 {
		query += " WHERE products.category_id =? AND products.sub_category_id =?"
		params = append(params, getCategoriesParams.CategoryID, getCategoriesParams.SubCategoryID)
	} else if getCategoriesParams.CategoryID == 0 && getCategoriesParams.SubCategoryID != 0 {
		query += " WHERE products.sub_category_id =?"
		params = append(params, getCategoriesParams.SubCategoryID)
	}

	if getCategoriesParams.CategoryID != 0 || getCategoriesParams.SubCategoryID != 0 {
		query += " AND products.name LIKE ?"
	} else {
		query += " WHERE products.name LIKE ?"
	}

	params = append(params, "%"+getCategoriesParams.TemporarySearchQuery+"%")

	if strings.ToLower(getCategoriesParams.TemporarySortQuery) == "desc" {
		query += " order by price desc"
	} else {
		query += " order by price asc"
	}

	query += " limit ? offset ?"

	params = append(params, getCategoriesParams.Size, getCategoriesParams.Page*getCategoriesParams.Size)

	if err := repo.db.Raw(query, params...).Find(&Products).Error; err != nil {
		return []models.Product{}
	}

	return Products
}

func (repo ProductRepo) GetProduct(id uint) (models.Product, error) {
	var product models.Product

	query := `select id, name, description, category_id, sub_category_id, icon, price, quantity, discount, status from products where id =?`

	if err := repo.db.Raw(query, id).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}
