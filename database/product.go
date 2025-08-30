package database

var productList[] Product

type Product struct{
	ID int             `json:"id"`
	Title string       `json:"name"`
	Description string `json:"description"`
	Price float64      `json:"price"`
	ImgUrl string      `json:"imageUrl"`
}


func Store(p Product) Product{
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product{
	return productList
}

func Get(productID int) *Product{
	for _, product := range productList{
		if product.ID == productID{
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for i, p := range productList{
		if p.ID == product.ID{
			productList[i] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product 
	for _, p := range productList{
		if p.ID != productID{
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init(){
	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "testy food and healthy food",
		Price: 100,
		ImgUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}

	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "testy food and healthy food",
		Price: 99,
		ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSYnc4owRRH-m0i5-4t-xyiRMvzvhu-QAVF_g&s",
	}
	prd3 := Product{
		ID: 3,
		Title: "Banana",
		Description: "testy food and healthy food",
		Price: 15,
		ImgUrl: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/banana-cavendish_0.png?itok=xIgYOIE_-9FKLRtCr",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
}