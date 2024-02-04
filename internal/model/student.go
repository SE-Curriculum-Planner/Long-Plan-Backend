package model

import "time"

type EngMajor string

const (
	CPE  EngMajor = "CPE"
	ISNE EngMajor = "ISNE"
)

type Student struct {
	StudentId   int        `gorm:"primaryKey;column:student_id;type:int"`
	Major       EngMajor   `gorm:"column:major;type:MAJOR;not null" json:"major"`
	LastUpdated *time.Time `gorm:"column:last_updated;type:timestamp;not null"`

	CourseEnrolled []CourseEnrolled `gorm:"foreignKey:CourseNo" json:"course_enrolled"`
}
