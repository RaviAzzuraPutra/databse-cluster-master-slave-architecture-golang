package models

import (
	"time"

	"gorm.io/datatypes"
)

type Cases struct {
	ID               *string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Case_Number      *string `json:"case_number"`
	Case_Title       *string `json:"case_title"`
	Case_Description *string `json:"case_description"`
	Incident_Date    datatypes.Date
	Location         *string   `json:"location"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`

	Suspects []Suspects `json:"suspects" gorm:"foreignKey:Case_ID"`
}
