package entity

import (
	"time"
	"gorm.io/gorm"
)

type EnrollmentSchedule struct {
	gorm.Model
	ScheduleDate 		time.Time
	SubjectType 		string
	StudyDurationHR    	string

	StudentID *uint
	Student   Student `gorm:"foriegnKey:StudentID"`

	TeacherID *uint
	Teacher   Teacher `gorm:"foriegnKey:TeacherID"`

	SubjectID *uint
	Subject   Subject `gorm:"foriegnKey:SubjectID"`
}