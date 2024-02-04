package model

import "time"

type CourseEnrolled struct {
	CourseNo  int        `gorm:"primaryKey;type:int"`
	StudentId int        `gorm:"primaryKey;type:int"`
	Year      int        `gorm:"primaryKey;type:int"`
	Semester  int        `gorm:"primaryKey;type:int"`
	Grade     string     `gorm:"column:grade;type:varchar(2);not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`

	Course Course `gorm:"foreignKey:CourseNo" json:"course"`
}
