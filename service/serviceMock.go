package service

import (
	"userRepo/model"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct{
	mock.Mock
}

var _ UserRepService = &MockUserService{}

func (m *MockUserService) GetUserService()([]model.User, error){
	args:= m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func(m *MockUserService) GetUserById(id int) (*model.User, error){
	args:= m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
    return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(email string) (*model.User, error){
	args:=m.Called(email)
	if args.Get(0)== nil{
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}