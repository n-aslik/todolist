package service

import (
	"Todolist/models"
	"Todolist/package/repository"
	"fmt"
)

func AddTask(task models.Task) error {
	err := repository.InsertTasks(task)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func UpdateTask(task models.Task, id int) error {
	err := repository.EditTasks(task.Title, task.Description, task.UserID, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func DeleteTask(isdeleted bool, id int) error {
	err := repository.SoftDeleteTasks(isdeleted, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func CheckTaskasDone(isdone bool, id int) error {
	err := repository.CheckTasksasDone(isdone, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func PrintAllTasks(isdeleted, isdone, isblocked bool, id uint) (tasks []models.Task, err error) {
	tasks, err = repository.GetAllTasks(isdeleted, isdone, isblocked, id)
	if err != nil {

		return tasks, err
	}
	return tasks, nil
}
func PrintAllTasksByID(isdeleted, isblocked bool, tid, uid uint) (task []models.Task, err error) {
	task, err = repository.GetAllTasksByID(isdeleted, isblocked, tid, uid)
	if err != nil {

		return task, err
	}
	return task, nil
}
