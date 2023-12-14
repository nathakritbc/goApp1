package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "world"})
}

func todos(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{"message": "todo lists"})
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.GET("/todos", todos)
	router.Run(":9090")
}
