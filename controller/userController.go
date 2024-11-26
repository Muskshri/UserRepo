package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"userRepo/service"

	"github.com/gorilla/mux"
)

type UserController struct{
	services service.UserRepService
}

func NewUserController(sr service.UserRepService) *UserController{
       return &UserController{services: sr}
}

func (u *UserController)GetUsersController(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	user, err:= u.services.GetUserService()
	if err!= nil{
		http.Error(w, "Error fetching users...", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UserController)GetUserByIdController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    
	params:= mux.Vars(r)
	id, _:= strconv.Atoi(params["id"])
	user, err:= u.services.GetUserById(id)
	if err!= nil{
		http.Error(w, "User not found ...", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (u *UserController) GetUserByEmailController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Invalid email...", http.StatusBadRequest)
		return
	}

	user, err:= u.services.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}