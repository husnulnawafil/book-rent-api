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

type userRepository struct {
	Sql *gorm.DB
	Rds *redis.Client
}

func NewUserRepository(sql *gorm.DB, rds *redis.Client) *userRepository {
	return &userRepository{
		Sql: sql,
		Rds: rds,
	}
}

type UserRepositoriesInterface interface {
	Create(data *models.User) (user *models.User, err error)
	Get(id uint) (user *models.User, err error)
	Update(id uint, data interface{}) (err error)
	Delete(id uint) (err error)
	SetCache(f, id, data interface{}) (err error)
	GetCache(f, id, model interface{}) (data interface{}, err error)
	DeleteCache(f, id interface{}) (err error)
}

func (u *userRepository) Create(data *models.User) (user *models.User, err error) {
	tx := u.Sql.Create(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	user = data
	return
}

func (u *userRepository) Get(id uint) (user *models.User, err error) {
	tx := u.Sql.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (u *userRepository) Update(id uint, data interface{}) (err error) {
	user := &models.User{}
	tx := u.Sql.Model(&user).Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return errors.New("no_data_updated")
	}
	return
}

func (u *userRepository) Delete(id uint) (err error) {
	var (
		user = models.User{}
		book = models.Book{}
	)
	u.Sql.Transaction(func(tx *gorm.DB) error {
		if err = u.Sql.Where("owner = ?", id).Delete(&book).Error; err != nil {
			return err
		}

		if err = u.Sql.Where("id = ?", id).Delete(&user).Error; err != nil {
			return err
		}

		return nil
	})
	return
}

func (u *userRepository) SetCache(f, id, data interface{}) (err error) {
	c := context.Background()
	rdsData, err := json.Marshal(data)
	if err != nil {
		return
	}

	key := modules.GenerateRedisKey(f, id)

	err = u.Rds.Set(c, key, rdsData, 0).Err()

	return
}

func (u *userRepository) GetCache(f, id, model interface{}) (data interface{}, err error) {
	c := context.Background()
	key := modules.GenerateRedisKey(f, id)
	res, err := u.Rds.Get(c, key).Result()
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

func (u *userRepository) DeleteCache(f, id interface{}) (err error) {
	c := context.Background()
	key := modules.GenerateRedisKey(f, id)
	err = u.Rds.Del(c, key).Err()
	return
}
