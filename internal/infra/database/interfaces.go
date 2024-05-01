package database

import "github.com/vitorpcruz/goexpert/9-APIS/internal/entity"

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
