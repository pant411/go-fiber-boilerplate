package repository

import (
	"github.com/pant411/go-fiber-boilerplate/modules/todo/model"
	"gorm.io/gorm"
)

// TodoRepository handles CRUD operations for todos
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new TodoRepository instance
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

// GetAllTodos returns all todos
func (r *TodoRepository) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// CreateTodo creates a new todo
func (r *TodoRepository) CreateTodo(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

// GetTodoByID returns a todo by ID
func (r *TodoRepository) GetTodoByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// UpdateTodo updates an existing todo
func (r *TodoRepository) UpdateTodo(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

// DeleteTodo deletes a todo
func (r *TodoRepository) DeleteTodo(id uint) error {
	return r.db.Delete(&model.Todo{}, id).Error
}
