package service

import (
	"app/common/herrors"
	"app/model"
	"app/repository"
	"context"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	EditUser(ctx context.Context, id, name, password string) error
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Create(ctx context.Context, user *model.User) (*model.User, error) {

	existedUser, err := s.userRepository.FindByEmail(ctx, user.Email)

	if err != nil {
		return nil, err
	}

	if existedUser != nil {
		return nil, herrors.ErrExistingUser
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	id, err := s.userRepository.Insert(ctx, user)

	if err != nil {
		return nil, err
	}

	return s.userRepository.FindByID(ctx, id)
}

func (s *userService) EditUser(ctx context.Context, id, name, password string) error {

	existedUser, err := s.userRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if existedUser == nil {
		return herrors.ErrUserNotFound
	}

	user := &model.User{
		ID:       id,
		Name:     name,
		Password: password,
		Email:    existedUser.Email,
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	if err := s.userRepository.UpdateOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	return s.userRepository.FindAll(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, herrors.ErrUserNotFound
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {

	existedUser, err := s.userRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if existedUser == nil {
		return herrors.ErrUserNotFound
	}

	err = s.userRepository.DeleteOne(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
