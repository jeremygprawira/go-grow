package repository

import (
	"go-grow/model"

	"gorm.io/gorm"
)

type BaseRepository interface {
    // User Account Management
    CreateUserToDB(user *model.User) (*model.User, error)
    FindUserByEmail(email string) (*model.User, error)
    UpdateUserToDB(user *model.User) (*model.User, error)

    CreateCoolToDB(cool *model.Cool) (*model.Cool, error)
}

type baseRepository struct {
    db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *baseRepository {
    return &baseRepository{db}
}

func (r *baseRepository) CreateUserToDB(user *model.User) (*model.User, error) {
    err := r.db.Create(&user).Error
    if err != nil {
        return user, err
    }

    return user, nil
}

func (r *baseRepository) FindUserByEmail(email string) (*model.User, error) {
    var user *model.User
    err := r.db.Where("email = ?", email).Find(&user).Error
    if err != nil {
        return user, err
    }

    return user, nil
}

func (r *baseRepository) UpdateUserToDB(user *model.User) (*model.User, error) {
    err := r.db.Save(&user).Error
    if err != nil {
        return user, err
    }

    return user, nil
}

func (r *baseRepository) CreateCoolToDB(cool *model.Cool) (*model.Cool, error) {
    err := r.db.Create(&cool).Error
    if err != nil {
        return cool, err
    }

    return cool, nil
}