package service

import (
	"userRepo/model"
	"userRepo/repository"
)

type UserRepService interface{
	GetUserService() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type userService struct{
	repo repository.UserRepo
}

func NewUserService(repository repository.UserRepo) UserRepService{
	 return &userService{repo: repository}
}

func (u *userService) GetUserService() ([]model.User, error){
    return u.repo.FindAll()
}

func(u *userService) GetUserById(id int) (*model.User, error){
	return u.repo.FindById(id)
}

func (u *userService) GetUserByEmail(email string) (*model.User, error){
	return u.repo.FindByEmail(email)
}
