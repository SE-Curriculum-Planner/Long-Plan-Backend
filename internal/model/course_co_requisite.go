package model

import "time"

type CourseCoRequisite struct {
	CourseNo      int        `gorm:"primaryKey;type:course_no"`
	CoRequisiteNo int        `gorm:"primaryKey;type:co_requisite_no"`
	CreatedAt     *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
