package main

import (
	"fmt"
	"log"
	"net/http"
	"userRepo/controller"
	"userRepo/repository"
	"userRepo/service"

	"github.com/gorilla/mux"
)

func main(){
	userRepository:= repository.NewRepo()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
    router:= mux.NewRouter()
	router.HandleFunc("/api/users", userController.GetUsersController).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.GetUserByIdController).Methods("GET")
	router.HandleFunc("/api/user/email",userController.GetUserByEmailController).Methods("GET")
    

	fmt.Println("Staring server on :8085")
	log.Fatal(http.ListenAndServe(":8085", router))
}