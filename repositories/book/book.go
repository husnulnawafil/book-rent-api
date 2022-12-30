package repositories

import (
	"github.com/go-redis/redis/v9"
	"github.com/husnulnawafil/dot-id-task/models"
	"github.com/husnulnawafil/dot-id-task/modules"
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
	Get(id uint) (book *models.Book, err error)
	List(pagination *modules.Pagination) (books []*models.Book, pgn *modules.Pagination, err error)
	Update(id uint, data interface{}) (err error)
	Delete(id uint) (err error)
	ListByOwner(owner uint) (books []*models.Book, err error)
}

func (b *bookRepository) Create(data *models.Book) (book *models.Book, err error) {
	tx := b.Sql.Create(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	book = data
	return
}

func (b *bookRepository) Get(id uint) (book *models.Book, err error) {
	tx := b.Sql.First(&book, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (b *bookRepository) ListByOwner(owner uint) (books []*models.Book, err error) {
	tx := b.Sql.Where("owner = ?", owner).Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (b *bookRepository) List(pagination *modules.Pagination) (books []*models.Book, pgn *modules.Pagination, err error) {
	paginate := modules.Paginate(books, pagination, b.Sql)
	tx := b.Sql.Scopes(paginate).Find(&books)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}

	pgn = pagination
	return books, pgn, err
}

func (b *bookRepository) Update(id uint, data interface{}) (err error) {
	book := &models.Book{}
	tx := b.Sql.Model(&book).Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return err
	}
	return
}

func (b *bookRepository) Delete(id uint) (err error) {
	var book models.Book
	tx := b.Sql.Where("id = ?", id).Delete(&book)
	if tx.Error != nil {
		return tx.Error
	}

	return
}
