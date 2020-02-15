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
	c.Header("Content-Type", "application/json")
	taskId := c.Param("id")
	var task models.Task
	err := models.GetParticularTask(&task, taskId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func CreateNewTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var task *models.Task
	c.BindJSON(&task)
	err := models.CreateTask(task)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"id": taskId,
	})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	taskId := c.Param("id")
	err := models.DeleteTask(&task, taskId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": taskId,
		})
	}
	// c.Header("Content-Type", "application/json")
}
