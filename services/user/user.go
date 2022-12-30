package services

import (
	"encoding/json"
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
	Update(id uint, data interface{}) (user *models.User, code int, err error)
	Delete(id uint) (err error)
}

func (u *userService) Create(data *models.User) (user *models.User, code int, err error) {
	user, err = u.userRepo.Create(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return
}

func (u *userService) Get(id uint) (user *models.User, code int, err error) {
	res, _ := u.userRepo.GetCache(u.userRepo.Get, id, user)
	if res != nil {
		byteData, err := json.Marshal(res)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		err = json.Unmarshal(byteData, &user)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return user, http.StatusOK, nil
	}

	user, err = u.userRepo.Get(id)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, errors.New("user_not_found")
	}

	books, _ := u.bookRepo.ListByOwner(id)
	if len(books) > 0 {
		user.Books = books
	}

	if err = u.userRepo.SetCache(u.userRepo.Get, id, user); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return
}

func (u *userService) Update(id uint, data interface{}) (user *models.User, code int, err error) {
	u.userRepo.DeleteCache(u.userRepo.Get, id)
	err = u.userRepo.Update(id, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	user, err = u.userRepo.Get(id)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, errors.New("user_not_found")
	}

	return
}

func (u *userService) Delete(id uint) (err error) {
	u.userRepo.DeleteCache(u.userRepo.Get, id)
	err = u.userRepo.Delete(id)
	return
}
