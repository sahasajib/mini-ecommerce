package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		http.Error(w, "Please provide a valid ID", http.StatusBadRequest)
		return
	}

	
	

	database.Delete(id)

	util.SendData(w, "Successfully product deleted", 201)
}