package repositories

import (
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
	Update(id uint, data *models.User) (user *models.User, err error)
	Delete(id uint) (user *models.User, err error)
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

func (u *userRepository) Update(id uint, data *models.User) (user *models.User, err error) {
	return
}

func (u *userRepository) Delete(id uint) (user *models.User, err error) {
	return
}
