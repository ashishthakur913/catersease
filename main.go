package main

import (
	"log"

	"github.com/ashishthakur913/CaterEase/api"
	"github.com/ashishthakur913/CaterEase/models"
	"github.com/ashishthakur913/CaterEase/repositories"
	"github.com/ashishthakur913/CaterEase/services"
	"github.com/ashishthakur913/CaterEase/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    e := echo.New()

    // Enable HTTP access logging
    e.Use(middleware.Logger())

    dsn := "host=db port=5432 user=user password=password dbname=mydb sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("error connecting to the database: %s", err)
    }

    // Migrate the schema
    err = db.AutoMigrate(&models.User{}, &models.Order{}, &models.Driver{})
    if err != nil {
        log.Fatalf("error migrating the schema: %s", err)
	}

	userService := services.NewUserService(repositories.NewUserRepository(db))
	userController := controllers.NewUserController(userService)

	handlers := &controllers.Handlers{
		UserController:  userController,
	}

	api.RegisterHandlers(e, handlers)

	e.Logger.Fatal(e.Start(":8080"))
}
