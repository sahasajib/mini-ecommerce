package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)


func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		http.Error(w, "Please provide a valid ID", http.StatusBadRequest)
		return
	}
	product := database.Get(id)
	if product == nil{
		util.SendError(w, http.StatusNotFound, "Product not found")
	}
	util.SendData(w, product, http.StatusNotFound)
}