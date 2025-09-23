package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)


func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		util.SendError(w, http.StatusBadRequest, "Plz give me valid if")
		return
	}
	product, err := h.svc.Get(id)
	if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if product == nil{
		util.SendError(w, http.StatusNotFound, "Product not found")
	}
	util.SendData(w,http.StatusNotFound ,product)
}