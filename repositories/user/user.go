package repositories

import (
	"errors"

	"github.com/go-redis/redis/v9"
	"github.com/husnulnawafil/dot-id-task/models"
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
