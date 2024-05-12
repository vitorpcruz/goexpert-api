package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/dto"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/entity"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/infra/database"
)

type UserHandler struct {
	UserRepositoryInterface database.UserRepositoryInterface
}

func NewUserHandler(
	userRepo database.UserRepositoryInterface,
) *UserHandler {
	return &UserHandler{
		UserRepositoryInterface: userRepo,
	}
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
		log.Printf("An error occurred while user was created: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var userJwtInput dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&userJwtInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserRepositoryInterface.FindByEmail(userJwtInput.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(userJwtInput.Password) {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	}

	_, tokenString, _ := jwt.Encode(claims)

	accessToken := struct {
		AccessToken string `json:"acess_token"`
	}{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
