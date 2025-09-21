package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)


func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	createdUser := newUser.Store()
	util.SendData(w, createdUser, http.StatusCreated)
}