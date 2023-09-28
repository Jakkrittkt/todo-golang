package model

import (
	"time"

	"github.com/google/uuid"
)

type TodoStatus string

const (
	IN_PROGRESS TodoStatus = "IN_PROGRESS"
	COMPLETED   TodoStatus = "COMPLETED"
)

type Todo struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;" json:"id"`
	Title       string     `gorm:"not null;" json:"title" `
	Description string     `json:"description"`
	Date        time.Time  `gorm:"autoCreateTime" json:"date" `
	Image       string     `json:"image"`
	Status      TodoStatus `gorm:"not null" json:"status" `
}

func (Todo) TableName() string {
	return "todos"
}
