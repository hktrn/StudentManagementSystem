package model

import (
	"time"
)

type Student struct {
	StudentID     uint32 `json:"student_id" gorm:"primaryKey;unique"`
	Name          string `json:"name"`
	Program       string `json:"program"`
	Major         string `json:"major"`
	AdmissionYear uint32 `json:"admission_year"`
	Coop          string `json:"coop" gorm:"default:'no"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Placement struct {
	StudentID uint32 `json:"student_id" gorm:"primaryKey;unique"`
	Company   string `json:"company"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	Term      string `json:"term"`
}

type Record struct {
	StudentID uint32    `json:"student_id" gorm:"primaryKey;unique"`
	Student   Student   `gorm:"foreignKey:StudentID"`
	Placement Placement `gorm:"foreignKey:StudentID"`
}
