package models

import (
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	TopicID          uint      `json:"topic_id" gorm:"not null; autoIncrement"`
	Name             string    `json:"name"`
	Description      string    `json:"description,omitempty"`
	CurrentInterval  int       `json:"current_interval"`
	NextRevisionDate time.Time `json:"next_revision_date"`
}

type RevisionLog struct {
	gorm.Model
	TopicID  uint   `json:"topic_id"`
	Interval int    `json:"interval"`
	Status   string `json:"status"`
}
