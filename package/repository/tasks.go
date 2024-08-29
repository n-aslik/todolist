package repository

import (
	"Todolist/database"
	"Todolist/logger"
	"Todolist/models"
)

func InsertTasks(task models.Task) error {
	err := database.GetconnectDB().Create(&task).Error
	if err != nil {
		logger.Error.Printf("[service.addtask]error in added task %s\n", err.Error())

	}
	return nil
}
func EditTasks(title, description string, userid, id int) error {
	err := database.GetconnectDB().Save(&models.Task{ID: id, Title: title, Description: description, UserID: userid}).Error
	if err != nil {
		logger.Error.Printf("[service.updatetask]error in update task %s\n", err.Error())
	}
	return nil
}

func SoftDeleteTasks(isdeleted bool, id int) error {
	err := database.GetconnectDB().Model(&models.Task{}).Where("id=?", id).Update("is_deleted", isdeleted).Error
	if err != nil {
		logger.Error.Printf("[service.delete task]error in deleted task %s\n", err.Error())
	}
	return nil
}

func GetAllTasks(isdeleted, isdone, isblocked bool, id uint) (task []models.Task, err error) {
	err = database.GetconnectDB().Preload("User").Joins("Join users ON users.id=tasks.user_id").Where("tasks.is_deleted=? AND tasks.is_done=?", isdeleted, isdone).Where("tasks.user_id=?", id).Order("tasks.id").Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeleted).Find(&task).Error
	if err != nil {
		logger.Error.Printf("[service.getalltasks]error in getting all task %s\n", err.Error())
		return task, translateErrors(err)
	}
	return task, nil
}
func GetAllTasksByID(isdeleted, isblocked bool, tid, uid uint) (task []models.Task, err error) {
	err = database.GetconnectDB().Preload("User").Joins("Join users ON users.id=tasks.user_id").Where("tasks.is_deleted=?", isdeleted).Where("tasks.id=? AND tasks.user_id=?", tid, uid).Where("users.is_blocked=? AND users.is_deleted=?", isblocked, isdeleted).Find(&task).Error
	if err != nil {
		logger.Error.Printf("[service.getalltasksbyid]error in getting all task by id %s\n", err.Error())
		return task, translateErrors(err)
	}
	return task, nil
}

func CheckTasksasDone(isdone bool, id int) error {
	var task models.Task
	err := database.GetconnectDB().Model(&task).Select("is_done").Where("id=?", id).Updates(models.Task{IsDone: isdone}).Error
	if err != nil {
		logger.Error.Printf("[service.checktaskasdone]error in checked task %s\n", err.Error())

	}
	return nil

}
