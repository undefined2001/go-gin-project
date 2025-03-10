package models

import "gorm.io/gorm"

// Task represents a task with a title and status
type Task struct {
	gorm.Model       // This includes ID, CreatedAt, UpdatedAt, DeletedAt
	Title  string    `json:"title"`                  // Title with JSON tag
	Status string    `json:"status"`                 // Status with JSON tag
}
