package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/entity"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbSqlite = "file:memory.db"

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(dbSqlite), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Jhon Doe", "j@j.com", "123456")
	userRepo := database.NewUserRepository(db)

	err = userRepo.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)

	assert.NotEmpty(t, user.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(dbSqlite), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	userRepo := database.NewUserRepository(db)

	err = userRepo.Create(user)
	assert.Nil(t, err)

	userFound, err := userRepo.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
