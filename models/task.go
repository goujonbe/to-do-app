package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

func GetAllTasks() []*Task {

	tasks := make([]*Task, 0)
	err := GetDB().Table("task").Find(&tasks).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tasks
}
