package services

import (
	"github.com/RobertMaulana/bookstore_users_api_microservice/domain/users"
	"github.com/RobertMaulana/bookstore_users_api_microservice/utils/crypto_utils"
	"github.com/RobertMaulana/bookstore_users_api_microservice/utils/date_utils"
	"github.com/RobertMaulana/bookstore_users_api_microservice/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)
type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if saveErr := result.Get(); saveErr != nil {
		return nil, saveErr
	}
	return result, nil
}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}