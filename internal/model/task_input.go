package model

import "time"

type TaskInput struct {
	Name    string `json:"name" binding:"required"`
	Status  string `json:"status"`   // todo, in_progress, done, canceled
	DueDate string `json:"due_date"` // ISO 8601 строка
}

func (ti TaskInput) ToTask() *Task {
	var dueDatePtr *time.Time
	if ti.DueDate != "" {
		if parsed, err := time.Parse(time.RFC3339, ti.DueDate); err == nil {
			dueDatePtr = &parsed
		}
	}

	return &Task{
		Name:    ti.Name,
		Status:  ti.Status,
		DueDate: dueDatePtr,
	}
}
