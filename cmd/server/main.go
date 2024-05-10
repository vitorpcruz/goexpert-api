package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vitorpcruz/goexpert/9-APIS/configs"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/entity"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/infra/database"
	handlers "github.com/vitorpcruz/goexpert/9-APIS/internal/infra/webserver"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	_, err = configs.LoadConfig(dir)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productRepo := database.NewProductRepo(db)
	productHandler := handlers.NewProductHandler(productRepo)

  userRepo := database.NewUserRepository(db)
  userHander := handlers.NewUserHandler(userRepo)

	// product

	log.Println("Running at 8080.")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetAll)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHander.CreateUser);

	http.ListenAndServe(":8080", r)
}
