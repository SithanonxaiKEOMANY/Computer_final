package models

import "time"

type Classroom struct {
	ID        uint `gorm:"primaryKey"`
	Major     string
	ClassYear int
	ClassName int
	CreatedAt time.Time
	UpdatedAt time.Time
}
