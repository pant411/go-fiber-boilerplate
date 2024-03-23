package service

import (
	"go-fiber-boilerplate/modules/todo/model"
	"go-fiber-boilerplate/modules/todo/repository"
)

// TodoService handles business logic for todos
type TodoService struct {
	todoRepo *repository.TodoRepository
}

// NewTodoService creates a new TodoService instance
func NewTodoService(todoRepo *repository.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

// GetAllTodos returns all todos
func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	return s.todoRepo.GetAllTodos()
}

// CreateTodo creates a new todo
func (s *TodoService) CreateTodo(todo *model.Todo) error {
	return s.todoRepo.CreateTodo(todo)
}

// GetTodoByID returns a todo by ID
func (s *TodoService) GetTodoByID(id uint) (*model.Todo, error) {
	return s.todoRepo.GetTodoByID(id)
}

// UpdateTodo updates an existing todo
func (s *TodoService) UpdateTodo(todo *model.Todo) error {
	return s.todoRepo.UpdateTodo(todo)
}

func (s *TodoService) DeleteTodo(id uint) error {
	return s.todoRepo.DeleteTodo(id)
}
