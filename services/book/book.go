package services

import (
	"github.com/husnulnawafil/dot-id-task/models"
	repositories "github.com/husnulnawafil/dot-id-task/repositories/book"
)

type bookService struct {
	bookRepo repositories.BookRepositoriesInterface
}

func NewBookService(bookRepo repositories.BookRepositoriesInterface) BookServiceInterface {
	return &bookService{
		bookRepo: bookRepo,
	}
}

type BookServiceInterface interface {
	Create(data *models.Book) (book *models.Book, err error)
	Get(id int) (book *models.Book, err error)
	List() (book []*models.Book, err error)
	Update(id int, data *models.Book) (book *models.Book, err error)
	Delete(id int) (book *models.Book, err error)
}

func (b *bookService) Create(data *models.Book) (book *models.Book, err error) {
	return
}

func (b *bookService) Get(id int) (book *models.Book, err error) {
	return
}

func (b *bookService) List() (book []*models.Book, err error) {
	return
}

func (b *bookService) Update(id int, data *models.Book) (book *models.Book, err error) {
	return
}

func (b *bookService) Delete(id int) (book *models.Book, err error) {
	return
}
