package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "Plz give me valid json")
		return
	}
	user, err := h.userRepo.Create(repo.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	
	util.SendData(w,  http.StatusCreated, user)
}