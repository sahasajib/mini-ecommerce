package repo

type Product struct{
	ID int             `json:"id"`
	Title string       `json:"name"`
	Description string `json:"description"`
	Price float64      `json:"price"`
	ImgUrl string      `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(ProductID int) error
}

type productRepo struct{
	productList []*Product
}
//constructor function
func NewProductRepo() ProductRepo{
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList{
		if product.ID == productID{
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error){
	return r.productList, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for i, p := range r.productList{
		if p.ID == product.ID{
			r.productList[i] = &product
		}
	}
	return &product, nil
}
func (r *productRepo) Delete(ProductID int) error {
	var tempList []*Product
	for _, p := range r.productList{
		if p.ID != ProductID{
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID: 1,
		Title: "Orange",
		Description: "testy food and healthy food",
		Price: 100,
		ImgUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}

	prd2 := &Product{
		ID: 2,
		Title: "Apple",
		Description: "testy food and healthy food",
		Price: 99,
		ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSYnc4owRRH-m0i5-4t-xyiRMvzvhu-QAVF_g&s",
	}
	prd3 := &Product{
		ID: 3,
		Title: "Banana",
		Description: "testy food and healthy food",
		Price: 15,
		ImgUrl: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/banana-cavendish_0.png?itok=xIgYOIE_-9FKLRtCr",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
}