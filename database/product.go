package database

var ProductList[] Product

type Product struct{
	ID int             `json:"id"`
	Title string       `json:"name"`
	Description string `json:"description"`
	Price float64      `json:"price"`
	ImgUrl string      `json:"imageUrl"`
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

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
}