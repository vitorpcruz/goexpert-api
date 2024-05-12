package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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

	configuration, err := configs.LoadConfig(dir)
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

	log.Println("Running at 8080.")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configuration.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configuration.JwtExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configuration.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetAll)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHander.CreateUser)
	r.Post("/users/token", userHander.GetJWT)

	http.ListenAndServe(":8080", r)
}
