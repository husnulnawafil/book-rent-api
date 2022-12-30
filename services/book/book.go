package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/husnulnawafil/dot-id-task/models"
	"github.com/husnulnawafil/dot-id-task/modules"
	bookRepositories "github.com/husnulnawafil/dot-id-task/repositories/book"
	userRepositories "github.com/husnulnawafil/dot-id-task/repositories/user"
)

type bookService struct {
	bookRepo bookRepositories.BookRepositoriesInterface
	userRepo userRepositories.UserRepositoriesInterface
}

func NewBookService(bookRepo bookRepositories.BookRepositoriesInterface, userRepo userRepositories.UserRepositoriesInterface) BookServiceInterface {
	return &bookService{
		bookRepo: bookRepo,
		userRepo: userRepo,
	}
}

type BookServiceInterface interface {
	Create(data *models.Book) (book *models.Book, code int, err error)
	Get(id uint) (book *models.Book, code int, err error)
	List(pagination *modules.Pagination) (books []*models.Book, pgn *modules.Pagination, code int, err error)
	Update(id uint, data interface{}) (book *models.Book, code int, err error)
	Delete(id uint) (err error)
}

func (b *bookService) Create(data *models.Book) (book *models.Book, code int, err error) {
	ownerID := data.Owner
	if _, err = b.userRepo.Get(ownerID); err != nil {
		return nil, http.StatusInternalServerError, errors.New("owner_not_found")
	}

	book, err = b.bookRepo.Create(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return
}

func (b *bookService) Get(id uint) (book *models.Book, code int, err error) {
	res, _ := b.bookRepo.GetCache(b.bookRepo.Get, id, book)
	if res != nil {
		byteData, err := json.Marshal(res)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		err = json.Unmarshal(byteData, &book)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return book, http.StatusOK, nil
	}

	book, err = b.bookRepo.Get(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = b.bookRepo.SetCache(b.bookRepo.Get, id, book); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return
}

func (b *bookService) List(pagination *modules.Pagination) (books []*models.Book, pgn *modules.Pagination, code int, err error) {
	type listBookRds struct {
		Books []*models.Book
		Pgn   *modules.Pagination
	}

	rdsData := &listBookRds{}
	res, _ := b.bookRepo.GetCache(b.bookRepo.List, nil, books)
	if res != nil {
		byteData, err := json.Marshal(res)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		err = json.Unmarshal(byteData, &rdsData)

		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		if rdsData.Pgn == pagination {
			return rdsData.Books, rdsData.Pgn, http.StatusOK, nil
		}

	}

	books, pgn, err = b.bookRepo.List(pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	rdsData = &listBookRds{
		Books: books,
		Pgn:   pgn,
	}

	if err = b.bookRepo.SetCache(b.bookRepo.List, nil, rdsData); err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return
}

func (b *bookService) Update(id uint, data interface{}) (book *models.Book, code int, err error) {
	b.userRepo.DeleteCache(b.bookRepo.Get, id)
	b.userRepo.DeleteCache(b.bookRepo.List, nil)
	err = b.bookRepo.Update(id, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	book, err = b.bookRepo.Get(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return
}

func (b *bookService) Delete(id uint) (err error) {
	book, _ := b.bookRepo.Get(id)
	if book == nil {
		return errors.New("book_not_found")
	}
	b.userRepo.DeleteCache(b.bookRepo.Get, id)
	b.userRepo.DeleteCache(b.bookRepo.List, nil)
	err = b.bookRepo.Delete(id)
	return
}
