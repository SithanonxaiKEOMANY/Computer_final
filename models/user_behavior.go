package models

import "time"

type UserBehavior struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	UserClassroomID  uint
	StudentCheck     bool
	StudentAbsent    bool
	StudentVacation  bool
	StudentBreakRule bool
	CountCheck       int
	CountAbsent      int
	CountVacation    int
	CountBreakRule   int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	User             User          `json:"user" gorm:"foreignKey:UserID"`
	UserClassroom    UserClassroom `json:"user_classroom" gorm:"foreignKey:UserClassroomID"`
}
