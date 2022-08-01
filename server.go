package main

import (
	"goapp/handler"
	"goapp/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repository.SetDatabase()
	defer repository.CloseDatabase()
	router.GET("/tasks", handler.GetTasks)
	router.PUT("/task", handler.InsertTask)
	router.PATCH("/task", handler.UpdateTask)
	router.Run(":4000")
}