package models

import "time"

type UserClassroom struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	ClassroomID uint
	SubjectID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User      `gorm:"foreignKey:UserID"`
	Classroom   Classroom `gorm:"foreignKey:ClassroomID"`
	Subject     Subject   `gorm:"foreignKey:SubjectID"`
}
