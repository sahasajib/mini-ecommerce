package repo

import (
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)



type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct{
	db *sqlx.DB
}
//constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo{
	return &productRepo{
		db: db,
	}


}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	INSERT INTO products (
		title,
		description,
		price,
		img_url
	) VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	r.db.QueryRow(query,
		p.Title,
		p.Description,
		p.Price,
		p.ImgUrl,
	).Scan(&p.ID)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	var product domain.Product
	query := `
	SELECT
		id,
		title,
		description,
		price,
		img_url
	FROM products
	WHERE id = $1
	`
	err := r.db.Get(&product, query, productID)
	if err != nil{
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) List() ([]*domain.Product, error){
	var productList []*domain.Product
	query := `
	SELECT
		id,
		title,
		description,
		price,
		img_url
	FROM products
	ORDER BY id DESC
	`
	err := r.db.Select(&productList, query)
	if err != nil{
		return nil, err
	}
	return productList, nil
}

func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
	UPDATE products SET
		title=$1,
		description=$2,
		price=$3,
		img_url=$4,
		updated_at=NOW()
	WHERE id=$5
	`
	row := r.db.QueryRow(query,
		product.Title,
		product.Description,
		product.Price,
		product.ImgUrl,
		product.ID,
	)
	err := row.Err()
	if err != nil{
		return nil, err
	}
	return &product, nil
}
func (r *productRepo) Delete(ProductID int) error {
	query := `
	DELETE FROM products
	WHERE id = $1
	`
	_, err := r.db.Exec(query, ProductID)
	if err != nil{
		return err
	}
	return nil
	
}

