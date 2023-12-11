package repository

import (
	"go-chi-api/entity"

	"github.com/go-pg/pg"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindUserByEmail(user entity.User) (entity.User, error)
	FindUserByID(userId uint) (entity.User, error)
	UpdateUser(user entity.User, userId uint) (entity.User, error)
	DeleteUser(userId uint) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *pg.DB) UserRepository {
	//return &UserRepositoryImpl{DB: db}
	return nil
}

func (r *UserRepositoryImpl) Create(user entity.User) (entity.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindUserByEmail(user entity.User) (entity.User, error) {
	err := r.DB.Model(&user).Where("email = ?", user.Email).First(&user).Error

	return user, err
}

func (r *UserRepositoryImpl) FindUserByID(userId uint) (entity.User, error) {
	user := entity.User{}
	err := r.DB.Model(&user).Where("id = ?", userId).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) UpdateUser(user entity.User, userId uint) (entity.User, error) {
	err := r.DB.Model(&user).Where("id = ?", userId).Updates(&user).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) DeleteUser(userId uint) error {
	user := entity.User{}
	err := r.DB.Model(&user).Delete(&user, userId).Error
	return err
}
