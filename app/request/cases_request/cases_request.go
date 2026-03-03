package cases_request

import (
	"databse-cluster-master-slave-architecture-golang/app/models"
	"time"
)

type Cases_Request struct {
	Case_Title       *string `form:"case_title"`
	Case_Description *string `form:"case_description"`
	Incident_Date    *string `form:"incident_date"`
	Location         *string `form:"location"`
}

type Cases_Dto struct {
	Case_Title       *string
	Case_Description *string
	Incident_Date    time.Time
	Location         *string
}

type Cases_Response struct {
	Case_Number      *string           `json:"case_number"`
	Case_Title       *string           `json:"case_title"`
	Case_Description *string           `json:"case_description"`
	Incident_Date    time.Time         `json:"incident_date"`
	Location         *string           `json:"location"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	Suspects         []models.Suspects `json:"suspects"`
}
