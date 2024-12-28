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

func (repo ProductRepo) GetProducts(getCategoriesParams types.GetCategoriesParams) ([]models.Product, int, error) {
	var products []models.Product
	var totalCount int

	// Start building the base query
	var query strings.Builder
	query.WriteString(`
		SELECT 
			products.id, 
			products.name, 
			products.description, 
			products.category_id, 
			products.sub_category_id, 
			products.icon, 
			products.price, 
			products.quantity, 
			products.discount, 
			products.status
		FROM products
	`)

	// Start building the count query
	var countQuery strings.Builder
	countQuery.WriteString(`
		SELECT COUNT(*) 
		FROM products
	`)

	// Store query parameters
	var params []interface{}

	// Apply filters
	whereClauses := []string{}
	if getCategoriesParams.CategoryID != 0 {
		whereClauses = append(whereClauses, "products.category_id = ?")
		params = append(params, getCategoriesParams.CategoryID)
	}
	if getCategoriesParams.SubCategoryID != 0 {
		whereClauses = append(whereClauses, "products.sub_category_id = ?")
		params = append(params, getCategoriesParams.SubCategoryID)
	}
	if getCategoriesParams.TemporarySearchQuery != "" {
		whereClauses = append(whereClauses, "products.name LIKE ?")
		params = append(params, "%"+getCategoriesParams.TemporarySearchQuery+"%")
	}

	// Append WHERE clause if there are any conditions
	if len(whereClauses) > 0 {
		whereCondition := " WHERE " + strings.Join(whereClauses, " AND ")
		query.WriteString(whereCondition)
		countQuery.WriteString(whereCondition)
	}

	// Apply sorting
	sortOrder := "ASC"
	if strings.ToLower(getCategoriesParams.TemporarySortQuery) == "desc" {
		sortOrder = "DESC"
	}
	query.WriteString(" ORDER BY products.price " + sortOrder)

	// Apply pagination
	query.WriteString(" LIMIT ? OFFSET ?")
	params = append(params, getCategoriesParams.Size, getCategoriesParams.Page*getCategoriesParams.Size)

	// Execute the count query to get the total count
	if err := repo.db.Raw(countQuery.String(), params[:len(params)-2]...).Scan(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Execute the main query to get the paginated results
	if err := repo.db.Raw(query.String(), params...).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

func (repo ProductRepo) GetProduct(id uint) (models.Product, error) {
	var product models.Product

	query := `select id, name, description, category_id, sub_category_id, icon, price, quantity, discount, status from products where id =?`

	if err := repo.db.Raw(query, id).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}
