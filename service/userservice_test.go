package service

import (
	"errors"
	"testing"
	"userRepo/model"
	"userRepo/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockRepo = new(repository.MockUserRepository)
	service = NewUserService(mockRepo)
)

func TestGetUsersServiceSuccess(t *testing.T){
	t.Run("GetUserService Success", func(t *testing.T){
		expectedUser:= []model.User{
			{
				Id:    1,
				Name:  "Muskan",
				Email: "muskan@gmail.com",
				Age:   25,
			},
		}
		mockRepo.On("FindAll").Return(expectedUser, nil)
		user, err:= service.GetUserService()

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)

	})
}

func TestGetUsersServiceError(t *testing.T){
	mockRepo:= new(repository.MockUserRepository)
	service = NewUserService(mockRepo)
	t.Run("GetUserService Error", func(t *testing.T){
		mockRepo.On("FindAll").Return([]model.User{}, errors.New("databse error"))
		user, err:= service.GetUserService()

		assert.Error(t, err)
		assert.Empty(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByIdServiceSuccess(t *testing.T){
	mockRepo:= new(repository.MockUserRepository)
	service:= NewUserService(mockRepo)
	t.Run("GetUserByIdService Success", func(t *testing.T){
		expectedUser:= &model.User{
				Id:    1,
				Name:  "Muskan",
				Email: "muskan@gmail.com",
				Age:   25,
		}
		mockRepo.On("FindById", 1).Return(expectedUser, nil)
		user, err:= service.GetUserById(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByIdServiceError(t *testing.T){
	mockRepo:= new(repository.MockUserRepository)
	service:= NewUserService(mockRepo)
	t.Run("GetUserById Error", func(t *testing.T){
		mockRepo.On("FindById",9999).Return(nil, gorm.ErrRecordNotFound)
		user, err:= service.GetUserById(9999)

		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmailServiceSuccess(t *testing.T){
	mockRepo:= new(repository.MockUserRepository)
	service:= NewUserService(mockRepo)
	t.Run("GetUserByEmail Success", func(t *testing.T){
		expectedUser:= &model.User{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
	    }
		mockRepo.On("FindByEmail", "muskan@gmail.com").Return(expectedUser, nil)
		user, err:= service.GetUserByEmail("muskan@gmail.com")

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmailServiceError(t *testing.T){
	mockRepo:= new(repository.MockUserRepository)
	service:= NewUserService(mockRepo)
	t.Run("GetUserByEmail Error", func(t *testing.T){
		mockRepo.On("FindByEmail", " ").Return(nil, gorm.ErrRecordNotFound)
		user, err:= service.GetUserByEmail(" ")

		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
    
}