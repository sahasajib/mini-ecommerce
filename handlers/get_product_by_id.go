package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)


func GetProductByID(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		http.Error(w, "Please provide a valid ID", http.StatusBadRequest)
		return
	}
	for _, product := range database.ProductList{
		if product.ID == id {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}
	util.SendData(w, "Product not found", http.StatusNotFound)
}