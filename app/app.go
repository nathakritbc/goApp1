package app

import (
	"database/sql"

	"go_app1/controllers"

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
	v1 := r.Group("/api/v1")

	// Simple group: v1
	// v1 := r.Group("/api/v1")
	controller := controllers.NewMangaController(a.DB)
	controllerTodo := controllers.NewTodoController(a.DB)

	mangaRoutes := v1.Group("/manga")
	mangaRoutes.POST("/", controller.InsertManga)
	mangaRoutes.GET("/", controller.GetAllManga)
	mangaRoutes.GET("/:id", controller.GetOneManga)
	mangaRoutes.PUT("/:id", controller.UpdateManga)
	mangaRoutes.DELETE("/:id", controller.DeleteManga)

	// todos routes
	todoRoutes := v1.Group("/todos")
	todoRoutes.POST("/", controllerTodo.InsertTodo)
	todoRoutes.GET("/", controllerTodo.GetAllTodo)
	todoRoutes.GET("/:id", controllerTodo.GetOneTodo)
	todoRoutes.PUT("/:id", controllerTodo.UpdateTodo)
	todoRoutes.DELETE("/:id", controllerTodo.DeleteTodo)

	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
