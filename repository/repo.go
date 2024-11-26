package repository

import (
	"userRepo/database"
	"userRepo/model"
	"fmt"
)

type UserRepo interface{
	FindAll()  ([]model.User, error)
	FindById(id int)  (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userRepo struct{}

func NewRepo() UserRepo{
	return &userRepo{}
}

func (u *userRepo)FindAll()([]model.User, error){
    var users []model.User
	fmt.Println("kya likhu?",database.GetDB())
	err := database.GetDB().Find(&users).Error
	return users, err
}

func (u *userRepo)FindById(id int)(*model.User, error){
	 var user model.User
	 err:= database.GetDB().First(&user,id).Error
	 return &user, err
}

func (u *userRepo)FindByEmail(email string)(*model.User, error){
	var user model.User
	err:= database.GetDB().Where("email = ?", email).First(&user).Error
	return &user, err
}
