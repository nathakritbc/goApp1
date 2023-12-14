package repositories

import (
	model "go_app1/models"
)

type TodoRepositoryInterface interface {
	InsertTodo(model.PostTodo) bool
	GetAllTodo() []model.Todo
	GetOneTodo(uint) model.Todo
	UpdateTodo(uint, model.PostTodo) model.Todo
	DeleteTodo(uint) bool
}
