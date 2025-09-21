package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type ReqUpdateProduct struct{
	Title string       `json:"name"`
	Description string `json:"description"`
	Price float64      `json:"price"`
	ImgUrl string      `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		util.SendError(w, http.StatusBadRequest, "Please provide a valid ID")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Plz give me valid json")
		return
	}
	
	//req.ID = id

	_, err = h.productRepo.Update(repo.Product{
		ID: id,
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImgUrl: req.ImgUrl,
	})
	if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w,http.StatusOK ,"Successfully product updated")
}