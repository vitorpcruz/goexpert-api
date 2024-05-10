package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/dto"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/entity"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/infra/database"
)

type UserHandler struct {
	UserRepositoryInterface database.UserRepositoryInterface
	Jwt                     *jwtauth.JWTAuth
	JwtExpiresIn            int
}

func NewUserHandler(userRepo database.UserRepositoryInterface) *UserHandler {
	return &UserHandler{UserRepositoryInterface: userRepo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.UserRepositoryInterface.Create(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
