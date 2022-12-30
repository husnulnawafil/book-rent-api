package repositories

import (
	"context"
	"encoding/json"
	"errors"

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
	SetCache(f, id, data interface{}) (err error)
	GetCache(f, id, model interface{}) (data interface{}, err error)
	DeleteCache(f, id interface{}) (err error)
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

func (b *bookRepository) SetCache(f, id, data interface{}) (err error) {
	c := context.Background()
	rdsData, err := json.Marshal(data)
	if err != nil {
		return
	}

	key := modules.GenerateRedisKey(f, id)
	err = b.Rds.Set(c, key, rdsData, 0).Err()

	return
}

func (b *bookRepository) GetCache(f, id, model interface{}) (data interface{}, err error) {
	c := context.Background()
	key := modules.GenerateRedisKey(f, id)
	res, err := b.Rds.Get(c, key).Result()
	if err != nil {
		return nil, err
	}

	if res == "" {
		return nil, errors.New("data_not_found")
	}

	if err = json.Unmarshal([]byte(res), &model); err != nil {
		return nil, err
	}

	data = model
	return
}

func (b *bookRepository) DeleteCache(f, id interface{}) (err error) {
	c := context.Background()
	key := modules.GenerateRedisKey(f, id)
	err = b.Rds.Del(c, key).Err()
	return
}
