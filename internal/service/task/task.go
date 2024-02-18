package task

import "test_task/internal/entities"

type queueInterface interface {
	GetListTasks() []entities.Task
	AddTask(entities.Task)
}
