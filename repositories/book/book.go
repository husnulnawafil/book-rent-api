package repositories

import (
	"github.com/go-redis/redis/v9"
	"github.com/husnulnawafil/dot-id-task/models"
	"gorm.io/gorm"
)

type bookRepository struct {
	Sql *gorm.DB
	Rds *redis.Client
}

func NewBookRepository(sql *gorm.DB, rds *redis.Client) *bookRepository {
	return &bookRepository{
		Sql: sql,
		Rds: rds,
	}
}

type BookRepositoriesInterface interface {
	Create(data *models.Book) (book *models.Book, err error)
	Get(id int) (book *models.Book, err error)
	List() (book []*models.Book, err error)
	Update(id int, data *models.Book) (book *models.Book, err error)
	Delete(id int) (book *models.Book, err error)
}

func (b *bookRepository) Create(data *models.Book) (book *models.Book, err error) {
	return
}

func (b *bookRepository) Get(id int) (book *models.Book, err error) {
	return
}

func (b *bookRepository) List() (book []*models.Book, err error) {
	return
}

func (b *bookRepository) Update(id int, data *models.Book) (book *models.Book, err error) {
	return
}

func (b *bookRepository) Delete(id int) (book *models.Book, err error) {
	return
}
