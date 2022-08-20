package service

import (
	"app/common/herrors"
	"app/model"
	"app/repository/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		qt := assert.New(t)

		user := &model.User{
			Email:    "test@gmail.com",
			Password: "123456",
			Name:     "test",
		}

		repo := new(mocks.UserRepository)
		repo.On("FindByEmail", mock.Anything, mock.Anything).Return(nil, nil)
		repo.On("Insert", mock.Anything, mock.Anything).Return("1", nil)
		repo.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)

		service := NewUserService(repo)

		result, err := service.Create(context.Background(), user)

		qt.NoError(err)
		qt.Equal(user, result)

	})

	t.Run("ExistingUser", func(t *testing.T) {
		qt := assert.New(t)

		user := &model.User{
			Email:    "test@test.com",
			Password: "123456",
			Name:     "test",
		}

		repo := new(mocks.UserRepository)
		repo.On("FindByEmail", mock.Anything, mock.Anything).Return(user, nil)

		service := NewUserService(repo)

		result, err := service.Create(context.Background(), user)

		qt.Error(err)
		qt.Equal(herrors.ErrExistingUser, err)
		qt.Nil(result)
	})

}
