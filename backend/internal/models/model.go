package models

import "time"

type Applicant struct {
    ID              string    `json:"id" db:"id"`
    NRIC            string    `json:"nric" db:"nric"`
    Name            string    `json:"name" db:"name"`
    EmploymentStatus *string   `json:"employment_status,omitempty" db:"employment_status"`
    Sex             string    `json:"sex" db:"sex"`
    DateOfBirth     time.Time `json:"date_of_birth" db:"date_of_birth"`
    Household       *string   `json:"household,omitempty" db:"household"`
}

type Application struct {
    ID            string     `json:"id" db:"id"`
    SchemeID      string     `json:"scheme_id" db:"scheme_id"`
    ApplicantID   string     `json:"applicant_id" db:"applicant_id"`
    Status        string     `json:"status" db:"status"`
    Disbursed     *string    `json:"disbursed,omitempty" db:"disbursed"`
    DisbursedDate *time.Time `json:"disbursed_date,omitempty" db:"disbursed_date"`
}

type Benefit struct {
    ID        string  `json:"id" db:"id"`
    Name      string  `json:"name" db:"name"`
    Amount    float64 `json:"amount" db:"amount"`
    SchemeID  string  `json:"scheme_id" db:"scheme_id"`
}
type Dependent struct {
    ID              string `json:"id" db:"id"`
    Name            string `json:"name" db:"name"`
    EmploymentStatus string `json:"employment_status" db:"employment_status"`
    Sex             string `json:"sex" db:"sex"`
    Relation        string `json:"relation" db:"relation"`
    DateOfBirth     string `json:"date_of_birth" db:"date_of_birth"`
    ParentID        string `json:"parent_id" db:"parent_id"`
}

type Scheme struct {
    ID       string  `json:"id" db:"id"`
    Name     string  `json:"name" db:"name"`
    Criteria *string `json:"criteria,omitempty" db:"criteria"`
    Benefits *string `json:"benefits,omitempty" db:"benefits"`
}