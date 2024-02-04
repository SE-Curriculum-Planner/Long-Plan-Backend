package model

import "time"

type GroupType string

const (
	Core          GroupType = "Core"
	MajorRequired GroupType = "Major Required"
	MajorElective GroupType = "Major Elective"
	LearnerPerson GroupType = "Learner Person"
	CoCreator     GroupType = "Co-Creator"
	ActiveCitizen GroupType = "Active Citizen"
	FreeElective  GroupType = "Free Elective"
)

type Course struct {
	CourseNo          int        `gorm:"primaryKey;column:course_no;type:int"`
	Title             string     `gorm:"column:title;type:varchar(255);not null"`
	RecommendSemester int        `gorm:"column:recommend_semester;type:int;not null"`
	RecommendYear     int        `gorm:"column:recommend_year;type:int;not null"`
	Credit            int        `gorm:"column:credit;type:int;not null"`
	GroupName         GroupType  `gorm:"column:group_name;type:GROUP_TYPE;not null"`
	CreatedAt         *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	PreRequisites []CoursePreRequisite `gorm:"foreignKey:CourseNo" json:"pre_requisites"`
	CoRequisites  []CourseCoRequisite  `gorm:"foreignKey:CourseNo" json:"co_requisites"`
}
