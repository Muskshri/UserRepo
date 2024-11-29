package repository 

import (
	"userRepo/model"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct{
	mock.Mock
}

var _ UserRepo = &MockUserRepository{}

func (m *MockUserRepository) FindAll() ([]model.User, error){
	result:= m.Called()
	return result.Get(0).([]model.User),result.Error(1)
}

func (m *MockUserRepository) FindById(id int) (*model.User, error){
	result:= m.Called(id)
	if result.Get(0) == nil{
		return nil, result.Error(1)
	}
	return result.Get(0).(*model.User), result.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*model.User, error){
	result:= m.Called(email)
	if result.Get(0) == nil{
		return nil, result.Error(1)
	}
	return result.Get(0).(*model.User), result.Error(1)
}