package repository

import (
	"go-fiber-boilerplate/modules/user/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new TodoRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAllTodos returns all todos
func (r *UserRepository) FindAllUser() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateTodo creates a new todo
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindOneUser(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindOneUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(id uint, user *model.User) error {
	return r.db.Where("id = ?", id).Updates(user).Error
}
