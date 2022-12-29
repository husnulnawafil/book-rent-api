package services

import (
	"errors"
	"net/http"

	"github.com/husnulnawafil/dot-id-task/models"
	repositories "github.com/husnulnawafil/dot-id-task/repositories/user"
)

type userService struct {
	userRepo repositories.UserRepositoriesInterface
}

func NewUserService(userRepo repositories.UserRepositoriesInterface) UserServiceInterface {
	return &userService{
		userRepo: userRepo,
	}
}

type UserServiceInterface interface {
	Create(data *models.User) (user *models.User, code int, err error)
	Get(id uint) (user *models.User, code int, err error)
	Update(id uint, data *models.User) (user *models.User, code int, err error)
	Delete(id uint) (user *models.User, err error)
}

func (u *userService) Create(data *models.User) (user *models.User, code int, err error) {
	user, err = u.userRepo.Create(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return
}

func (u *userService) Get(id uint) (user *models.User, code int, err error) {
	user, err = u.userRepo.Get(id)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, errors.New("user_not_found")
	}
	return
}

func (u *userService) Update(id uint, data *models.User) (user *models.User, code int, err error) {
	return
}

func (u *userService) Delete(id uint) (user *models.User, err error) {
	return
}
