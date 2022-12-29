package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/husnulnawafil/dot-id-task/configs"
	bh "github.com/husnulnawafil/dot-id-task/handlers/book"
	uh "github.com/husnulnawafil/dot-id-task/handlers/user"
	br "github.com/husnulnawafil/dot-id-task/repositories/book"
	ur "github.com/husnulnawafil/dot-id-task/repositories/user"
	"github.com/husnulnawafil/dot-id-task/routes"
	bs "github.com/husnulnawafil/dot-id-task/services/book"
	us "github.com/husnulnawafil/dot-id-task/services/user"
	"github.com/husnulnawafil/dot-id-task/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config := configs.GetConfig()
	sql := utils.InitSQL(config)
	rds := utils.InitRedis(config)

	userRepo := ur.NewUserRepository(sql, rds)
	userService := us.NewUserService(userRepo)
	userHandler := uh.NewUserHandler(userService)

	bookRepo := br.NewBookRepository(sql, rds)
	bookService := bs.NewBookService(bookRepo)
	bookHandler := bh.NewUserHandler(bookService)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	routes.UserPath(e, &userHandler)
	routes.BookPath(e, &bookHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
