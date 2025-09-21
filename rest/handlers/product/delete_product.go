package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		util.SendError(w, http.StatusBadRequest, "Plz give me valid id")
		return
	}

	
	
	err = h.productRepo.Delete(id)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusAccepted ,"Successfully product deleted")
}