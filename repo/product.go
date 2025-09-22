package repo

import "github.com/jmoiron/sqlx"

type Product struct{
	ID int             `json:"id" db:"id"`
	Title string       `json:"name" db:"title"`
	Description string `json:"description" db:"description"`
	Price float64      `json:"price" db:"price"`
	ImgUrl string      `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(ProductID int) error
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(productID int) (*Product, error) {
	var product Product
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

func (r *productRepo) List() ([]*Product, error){
	var productList []*Product
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

func (r *productRepo) Update(product Product) (*Product, error) {
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

// func generateInitialProducts(r *productRepo) {
// 	prd1 := &Product{
// 		ID: 1,
// 		Title: "Orange",
// 		Description: "testy food and healthy food",
// 		Price: 100,
// 		ImgUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
// 	}

// 	prd2 := &Product{
// 		ID: 2,
// 		Title: "Apple",
// 		Description: "testy food and healthy food",
// 		Price: 99,
// 		ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSYnc4owRRH-m0i5-4t-xyiRMvzvhu-QAVF_g&s",
// 	}
// 	prd3 := &Product{
// 		ID: 3,
// 		Title: "Banana",
// 		Description: "testy food and healthy food",
// 		Price: 15,
// 		ImgUrl: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/banana-cavendish_0.png?itok=xIgYOIE_-9FKLRtCr",
// 	}

// 	r.productList = append(r.productList, prd1)
// 	r.productList = append(r.productList, prd2)
// 	r.productList = append(r.productList, prd3)
// }