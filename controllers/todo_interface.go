package controllers

import "github.com/gin-gonic/gin"

type TodoControllerInterface interface {
	InsertTodo(*gin.Context)
	GetAllTodo(*gin.Context)
	GetOneTodo(*gin.Context)
	UpdateTodo(*gin.Context)
	DeleteTodo(*gin.Context)
}
