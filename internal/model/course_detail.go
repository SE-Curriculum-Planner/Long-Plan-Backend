package model

import "time"

type CurriculumDetail struct {
	CurriculumId   int        `gorm:"primaryKey;column:curriculum_id;type:int"`
	GroupName      GroupType  `gorm:"column:group_name;type:GROUP_TYPE;not null"`
	RequireCredits int        `gorm:"column:require_credits;type:int;not null"`
	CreatedAt      *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
