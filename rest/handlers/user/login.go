package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)
type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){
	var req ReqLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil{
		util.SendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if usr == nil{
		util.SendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	accessToken, err:= util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
	})
	if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusOK, accessToken)
}