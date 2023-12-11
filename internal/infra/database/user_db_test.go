package database

import (
	"testing"

	"github.com/rafaelsouzaribeiro/9-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	User, _ := entity.NewUser("Rafael", "rafel@gmail.com", "123456")
	UserDb := NewUser(db)

	err = UserDb.Create(User)
	assert.Nil(t, err)

	var userFound *entity.User
	err = db.First(&userFound, "id=?", User.Id).Error
	assert.Nil(t, err)
	assert.Equal(t, User.Id, userFound.Id)
	assert.Equal(t, User.Name, userFound.Name)
	assert.Equal(t, User.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	User, _ := entity.NewUser("Rafael", "rafel@gmail.com", "123456")
	UserDb := NewUser(db)

	err = UserDb.Create(User)
	assert.Nil(t, err)

	// Declare userFound as a pointer to entity.User
	var userFound *entity.User
	userFound, err = UserDb.FindByEmail(User.Email)
	assert.Nil(t, err)

	// Now you can access userFound as a pointer
	assert.Equal(t, User.Id, userFound.Id)
	assert.Equal(t, User.Name, userFound.Name)
	assert.Equal(t, User.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
