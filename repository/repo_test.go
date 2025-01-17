package repository

import (
	"errors"
	"testing"
	"userRepo/model"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFindAllUsersSucess(t *testing.T){
	mockRepo:= new(MockUserRepository)
	t.Run("FindAll Success", func(t *testing.T){
        expectedUser:= []model.User{
			{
				Id: 1,
				Name:  "Muskan",
				Email: "muskan@gmail.com",
				Age:   25,
		   },
	    }
		mockRepo.On("FindAll").Return(expectedUser, nil)
		users, err:= mockRepo.FindAll()

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, users)
		mockRepo.AssertExpectations(t)
	})
}

func TestFindAllErrors(t *testing.T){
	mockRepo:= new(MockUserRepository)
	t.Run("FindAll Error", func(t *testing.T) {
	mockRepo.On("FindAll").Return([]model.User{}, errors.New("database error"))
	users, err1:= mockRepo.FindAll()
	
	assert.Error(t, err1)
	assert.Empty(t, users)
	mockRepo.AssertExpectations(t)
  })
}

func TestFindByIdSuccess(t *testing.T){
	mockRepo:= new(MockUserRepository)
	t.Run("FindById Success", func(t *testing.T) {
	expectedUser := &model.User{
		Id:    1,
		Name:  "Muskan",
		Email: "muskan@gmail.com",
		Age:   25,
	}

	mockRepo.On("FindById", 1).Return(expectedUser, nil)
	user, err:= mockRepo.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
	
  })
}

func TestFindByIdError(t *testing.T){
	mockRepo:= new(MockUserRepository)
	t.Run("FindById Not Found", func(t *testing.T) {
		mockRepo.On("FindById", 999999).Return(nil, gorm.ErrRecordNotFound)
		user, err:= mockRepo.FindById(999999)

		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})

}

func TestFindByEmailSuccess(t *testing.T){
	mockRepo:= new(MockUserRepository)
	t.Run("FindByEmail Success", func(t *testing.T){
		expectedUser:= &model.User{
			Id: 1,
			Name:"Muskan",
			Email:"muskan@gmail.com",
			Age: 25,
		}
        mockRepo.On("FindByEmail", "muskan@gmail.com").Return(expectedUser, nil)
		user, err:= mockRepo.FindByEmail("muskan@gmail.com")


		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestFindByEmailError(t *testing.T){
	mockRepo:= new(MockUserRepository)
	t.Run("FindById Not Found", func(t *testing.T){
		mockRepo.On("FindByEmail", "muskan@gmail.com").Return(nil, gorm.ErrRecordNotFound)
        user, err:= mockRepo.FindByEmail("muskan@gmail.com")


		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}