package model

import "time"

type CoursePreRequisite struct {
	CourseNo       int        `gorm:"primaryKey;type:course_no"`
	PreRequisiteNo int        `gorm:"primaryKey;type:pre_requisite_no"`
	CreatedAt      *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
