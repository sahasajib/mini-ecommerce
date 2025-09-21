package user

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)
type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){
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

	cnf := config.GetConfig()

	accessToken, err:= util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
	})
	if err != nil{
		http.Error(w, "Could not create access token", http.StatusInternalServerError)
		return
	}
	util.SendData(w, accessToken, http.StatusOK)
}