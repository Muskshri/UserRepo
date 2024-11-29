package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T){
	tests := []struct{
		name string
		user User
		expected string
	}{
		{"Valid user", User{Name:"Muskan Shrivastava",Email: "musk@gmail.com", Age: 25 },""},
		{"Empty User Name", User{Name:"",Email: "musk@gmail.com", Age: 25 },"Invalid user name.."},
		{"Empty Email Name", User{Name:"Muskan Shrivastava", Email:"",Age: 25},"Invalid user email.."},
	}

	for _,testcase:= range tests{
		t.Run(testcase.name, func(t *testing.T){
		err:= testcase.user.Validate()
		if testcase.expected == ""{
			assert.NoError(t, err)
		}else{
			assert.EqualError(t, err,testcase.expected)
		}
	   })
	}
}
