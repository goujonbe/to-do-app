package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goujonbe/to-do-app/models"
)

func GetAllTasks(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data := models.GetAllTasks()
	c.JSON(http.StatusOK, data)
}

func GetParticularTask(c *gin.Context) {
	taskId := c.Param("id")
	// database query to get the task with the given id
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"id": taskId,
	})
}

func CreateNewTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet",
	})
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"id": taskId,
	})
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"id": taskId,
	})
}
