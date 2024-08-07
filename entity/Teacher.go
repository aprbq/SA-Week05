package entity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	FirstName 		string
	LastName 		string
	Email    		string
	PhoneNumber    	string

	EnrollmentSchedule []EnrollmentSchedule `gorm:"foreignKey:TeacherID"`

}