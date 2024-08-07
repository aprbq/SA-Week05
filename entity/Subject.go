package entity

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectName 		string
	SubjectType 		string
	TotalstudyHours   	uint

	CourseTypeID 		*uint
	CourseType   CourseType `gorm:"foriegnKey:CourseTypeID"`

	EnrollmentSchedule []EnrollmentSchedule `gorm:"foreignKey:SubjectID"`

}