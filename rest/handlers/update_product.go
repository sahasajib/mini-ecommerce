package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


func UpdateProduct(w http.ResponseWriter, r *http.Request){
	ID := r.PathValue("id")

	id, err := strconv.Atoi(ID)

	if err != nil{
		http.Error(w, "Please provide a valid ID", http.StatusBadRequest)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}
	
	newProduct.ID = id

	database.Update(newProduct)

	util.SendData(w, "Successfully product updated", 201)
}