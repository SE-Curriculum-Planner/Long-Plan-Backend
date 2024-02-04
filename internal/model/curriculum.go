package model

import "time"

type Curriculum struct {
	ID                  int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CurriculumProgram   string     `gorm:"column:curriculum_program;type:varchar(255);not null"`
	Year                int        `gorm:"column:year;type:int;not null"`
	IsCOOPPlan          bool       `gorm:"column:is_coop_plan;type:boolean;not null"`
	RequireCredits      int        `gorm:"column:require_credits;type:int;not null"`
	FreeElectiveCredits int        `gorm:"column:free_elective_credits;type:int;not null"`
	CreatedAt           *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           *time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	CurriculumDetails []CurriculumDetail `gorm:"foreignKey:CurriculumId" json:"curriculum_details"`
}
