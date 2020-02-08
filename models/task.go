package models

import (
	"fmt"
)

type Task struct {
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

func CreateTask(task *Task) (err error) {
	if err = GetDB().Create(task).Error; err != nil {
		return err
	}
	return nil
}
