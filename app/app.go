package app

import (
	"database/sql"

	controller "go_app1/controllers"
	"go_app1/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewMangaController(a.DB)
	controllerTodo := controller.NewTodoController(a.DB)
	r.POST("/manga", controller.InsertManga)
	r.GET("/manga", controller.GetAllManga)
	r.GET("/manga/:id", controller.GetOneManga)
	r.PUT("/manga/:id", controller.UpdateManga)
	r.DELETE("/manga/:id", controller.DeleteManga)

	// todos routes

	r.POST("/todos", controller.controllerTodo)
	// r.GET("/todos", controllerTodo.GetAllTodo)
	// r.GET("/todos/:id", controllerTodo.GetOneTodo)
	// r.PUT("/todos/:id", controllerTodo.UpdateTodo)
	// r.DELETE("/todos/:id", controllerTodo.DeleteTodo)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
