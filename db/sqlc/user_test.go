package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/modaniru/learning/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T){
	createRandUser(t)
}

func TestGetUserList(t *testing.T){
	for i := 0; i < 10; i++{
		createRandUser(t)
	}
	input := GetAllParams{
		Limit: 5,
		Offset: 5,
	}
	list, err := testQueries.GetAll(context.Background(), input)
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	assert.Len(t, list, 5)
	for _, u := range list{
		assert.NotEmpty(t, u)
	}
}

func TestUpdateUser(t *testing.T){
	user := createRandUser(t)
	update := UpdateUserParams{
		ID: user.ID,
		Name: utils.RandString(5),
		Surname: utils.RandString(5),
	}
	newUser, err := testQueries.UpdateUser(context.Background(), update)
	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)
	//compare
	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, update.Name, newUser.Name)
	assert.Equal(t, update.Surname, newUser.Surname)
	assert.Equal(t, user.Login, newUser.Login)
}

func TestDeleteUser(t *testing.T){
	user := createRandUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	assert.NoError(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, user)
}

func TestGetUser(t *testing.T){
	user := createRandUser(t)
	newUser, err := testQueries.GetUser(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)
	// compare
	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, user.Login, newUser.Login)
	assert.Equal(t, user.Name, newUser.Name)
	assert.Equal(t, user.Surname, newUser.Surname)
}

func randUser() *CreateUserParams{
	return &CreateUserParams{
		Login: utils.RandString(10),
		Name: utils.RandString(6),
		Surname: utils.RandString(6),
	}
}

func createRandUser(t *testing.T) User{
	userParams := randUser()
	user, err := testQueries.CreateUser(context.Background(), *userParams)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	// Compare
	assert.Equal(t, userParams.Name, user.Name)
	assert.Equal(t, userParams.Surname, user.Surname)
	assert.Equal(t, userParams.Login, user.Login)
	assert.NotZero(t, user.ID)
	return user
}