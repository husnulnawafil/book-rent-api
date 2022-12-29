package services

import (
	"errors"
	"net/http"

	"github.com/husnulnawafil/dot-id-task/models"
	bookRepositories "github.com/husnulnawafil/dot-id-task/repositories/book"
	userRepositories "github.com/husnulnawafil/dot-id-task/repositories/user"
)

type userService struct {
	userRepo userRepositories.UserRepositoriesInterface
	bookRepo bookRepositories.BookRepositoriesInterface
}

func NewUserService(userRepo userRepositories.UserRepositoriesInterface, bookRepo bookRepositories.BookRepositoriesInterface) UserServiceInterface {
	return &userService{
		userRepo: userRepo,
		bookRepo: bookRepo,
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

	books, _ := u.bookRepo.ListByOwner(id)
	if len(books) > 0 {
		user.Books = books
	}

	return
}

func (u *userService) Update(id uint, data *models.User) (user *models.User, code int, err error) {
	return
}

func (u *userService) Delete(id uint) (user *models.User, err error) {
	return
}
