package models

import (
	"fmt"
)

type Task struct {
	ID          int    `json:"id" gorm:"AUTO_INCREMENT"`
	Title       string `json:"title" gorm:"unique;not null"`
	Description string `json:"description"`
	Complete    bool   `json:"complete" gorm:"default:false"`
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

func GetParticularTask(task *Task, id string) (err error) {
	if err := GetDB().Where("ID = ?", id).Find(&task).Error; err != nil {
		return err
	}
	return nil
}

func CreateTask(task *Task) (err error) {
	if err := GetDB().Create(task).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(task *Task, id string) (err error) {
	if err := GetDB().Where("ID = ?", id).Delete(&task).Error; err != nil {
		return err
	}
	return nil
}
