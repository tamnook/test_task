package queue

import (
	"sort"
	"test_task/internal/entities"
)

type Queue struct {
	queue []entities.Task
}

func (q *Queue) GetListTasks() (tasks []entities.Task) {
	copy(tasks, q.queue)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].TaskTakeDatetime.Before(tasks[j].TaskTakeDatetime)
	})
	return tasks
}
