package serializer

import "SK-todo/model"

type Task struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"eat"`
	Description string `json:"description" example:"sleep"`
	Status      int    `json:"status" example:"0"`
	CreatedAt   int64  `json:"created_at"`
	StartTime   int64  `json:"start_time"`
	DueDate     int64  `json:"due_date"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Status:      item.Status,
		CreatedAt:   item.CreatedAt.Unix(),
		StartTime:   item.StartTime,
		DueDate:     item.DueDate,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
