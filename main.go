package main

import (
	"github.com/goujonbe/to-do-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1.GET("/tasks", handlers.GetAllTasks)
	v1.GET("/tasks/:id", handlers.GetParticularTask)
	v1.POST("/tasks", handlers.CreateNewTask)
	v1.PUT("/tasks/:id", handlers.UpdateTask)
	v1.DELETE("/tasks/:id", handlers.DeleteTask)
	r.Run() // listen and serve on 0.0.0.0:8080
}
