package domain

type Product struct{
	ID int             `json:"id" db:"id"`
	Title string       `json:"name" db:"title"`
	Description string `json:"description" db:"description"`
	Price float64      `json:"price" db:"price"`
	ImgUrl string      `json:"imageUrl" db:"img_url"`
}