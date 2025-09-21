package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)
type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request){
	var reqLogin ReqLogin
	if err := json.NewDecoder(r.Body).Decode(&reqLogin); err != nil {
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	usr:= database.FindById(reqLogin.Email, reqLogin.Password)
	if usr == nil{
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	util.SendData(w, usr, http.StatusOK)
}