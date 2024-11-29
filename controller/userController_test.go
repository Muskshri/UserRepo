package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepo/model"
	"userRepo/service"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersSuccess(t *testing.T){
	service:= new(service.MockUserService)
	t.Run("GetUsers Success", func(t *testing.T) {
		expectedUser := []model.User{{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		},
		}
		service.On("GetUserService").Return(expectedUser, nil)
		controller:= NewUserController(service)
		req,_:= http.NewRequest("GET","/api/users", nil)
		w:= httptest.NewRecorder() //store http response
		controller.GetUsersController(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		service.AssertExpectations(t)
	})
}

func TestGetUsersError(t *testing.T){
	service:= new(service.MockUserService)
	controller:= NewUserController(service)
	t.Run("GetUser Error", func(t *testing.T){
		service.On("GetUserService").Return([]model.User{}, assert.AnError)
		req,_:= http.NewRequest("GET","/api/users", nil)
		w:= httptest.NewRecorder()
		controller.GetUsersController(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetUserByIdSuccess(t *testing.T){
	service:= new(service.MockUserService)
	controller:= NewUserController(service)
	t.Run("GetUserById Success", func(t *testing.T){
	expectedUser := &model.User{
		Id:    1,
		Name:  "Muskan",
		Email: "muskan@gmail.com",
		Age:   25,
	}
	service.On("GetUserById", 1).Return(expectedUser, nil)
	req,_:= http.NewRequest("GET","/api/users/1", nil)
    req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
    controller.GetUserByIdController(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	service.AssertExpectations(t)
   })
}

func TestGetUserByIdError(t *testing.T){
	service:= new(service.MockUserService)
	controller:= NewUserController(service)
	t.Run("GetUserById Error", func(t *testing.T){
	 service.On("GetUserById", 9999).Return(nil, assert.AnError)
	 req, _ := http.NewRequest("GET", "/api/user/1", nil)
	 req = mux.SetURLVars(req, map[string]string{"id": "9999"})
	 w := httptest.NewRecorder()

	 controller.GetUserByIdController(w, req)

	 assert.Equal(t, http.StatusInternalServerError, w.Code)
	 service.AssertExpectations(t)
	})
}

func TestGetUserByEmailSuccess(t *testing.T){
	service:= new(service.MockUserService)
	controller:= NewUserController(service)
	t.Run("GetUserByEmail Success", func(t *testing.T){
		expectedUser := &model.User{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		}
		service.On("GetUserByEmail","muskan@gmail.com").Return(expectedUser, nil)
		req,_:= http.NewRequest("GET", "/api/user/email?email=muskan@gmail.com", nil)
		w := httptest.NewRecorder()
		controller.GetUserByEmailController(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		service.AssertExpectations(t)
	})
}

func TestGetUserByEmailError(t *testing.T){
	service:= new(service.MockUserService)
	controller:= NewUserController(service)
	t.Run("GetUserByEmail Error",func(t *testing.T){
		service.On("GetUserByEmail", "muskan@gmail.com").Return(nil, assert.AnError)
		req, _ := http.NewRequest("GET", "/api/user/email?email=muskan@gmail.com", nil)
		w := httptest.NewRecorder()

		controller.GetUserByEmailController(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		service.AssertExpectations(t)
	})
}