package model

import "time"

type Task struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"` // todo, in_progress, done, canceled
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DueDate   *time.Time `json:"due_date,omitempty"`
}
