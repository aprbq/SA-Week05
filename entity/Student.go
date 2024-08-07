package entity

import (
	"time"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName 		string
	LastName 		string
	DateOfBirth    	time.Time
	Email    		string
	PhoneNumber    	string

	// 1 student ลงได้หลาย EnrollmentSchedule
	EnrollmentSchedule []EnrollmentSchedule `gorm:"foreignKey:StudentID"`

}