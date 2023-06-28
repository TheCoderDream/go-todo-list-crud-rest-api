package task

import (
	"gorm.io/gorm"
)

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Task struct {
	gorm.Model
	ID          uint     `json:"id" gorm:"primaryKey"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     string   `json:"dueDate"`
	Priority    Priority `json:"priority"`
	Completed   bool     `json:"completed"`
	UserID      uint     `json:"userID"`
}
