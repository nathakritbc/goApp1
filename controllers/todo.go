package controllers

import (
	"database/sql"

	model "go_app1/models"
	repository "go_app1/repositories"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Db *sql.DB
}

func NewTodoController(db *sql.DB) TodoControllerInterface {
	return &TodoController{Db: db}
}

// DeleteTodo implements TodoControllerInterface
func (m *TodoController) DeleteTodo(c *gin.Context) {
	DB := m.Db
	var uri model.TodoUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTodoRepository(DB)
	// fmt.Println(uri.ID)

	delete := repository.DeleteTodo(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete todo successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete todo failed"})
		return
	}
}

// GetAllTodo implements TodoControllerInterface
func (m *TodoController) GetAllTodo(c *gin.Context) {
	DB := m.Db
	repository := repository.NewTodoRepository(DB)
	get := repository.GetAllTodo()

	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get todo successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "todo not found"})
		return
	}
}

// GetOneTodo implements TodoControllerInterface
func (m *TodoController) GetOneTodo(c *gin.Context) {
	DB := m.Db
	var uri model.TodoUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTodoRepository(DB)
	get := repository.GetOneTodo(uri.ID)
	if (get != model.Todo{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get todo successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "todo not found"})
		return
	}
}

// InsertTodo implements TodoControllerInterface
func (m *TodoController) InsertTodo(c *gin.Context) {
	DB := m.Db
	var post model.PostTodo
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTodoRepository(DB)
	insert := repository.InsertTodo(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert todo successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert todo failed"})
		return
	}
}

// UpdateTodo implements TodoControllerInterface
func (m *TodoController) UpdateTodo(c *gin.Context) {
	DB := m.Db
	var post model.PostTodo
	var uri model.TodoUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTodoRepository(DB)
	update := repository.UpdateTodo(uri.ID, post)
	if (update != model.Todo{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update todo successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update todo failed"})
		return
	}
}
