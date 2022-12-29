package utils

import (
	"fmt"

	"github.com/husnulnawafil/dot-id-task/configs"
	"github.com/husnulnawafil/dot-id-task/models"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitSQL(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.MySql.Username,
		config.MySql.Password,
		config.MySql.Address,
		config.MySql.Port,
		config.MySql.Name,
	)

	sql, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(sql)
	return sql
}

func InitialMigration(sql *gorm.DB) {
	sql.AutoMigrate(&models.User{})
	sql.AutoMigrate(&models.Book{})
}
