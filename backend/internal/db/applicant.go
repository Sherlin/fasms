package db

import (

    log "github.com/sirupsen/logrus"

	"fasms/internal/models"
)
func CreateApplicant(newApplicant models.Applicant) error {
	// Prepare the SQL query to insert a new applicant
	query := `
		INSERT INTO applicants (id, nric, name, employment_status, sex, date_of_birth, household) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Execute the query with the provided values
	_, err := DB.Exec(query, newApplicant.ID, newApplicant.NRIC, newApplicant.Name, newApplicant.EmploymentStatus, 
		newApplicant.Sex, newApplicant.DateOfBirth, newApplicant.Household)
	
	if err != nil {
		log.Error("Error inserting applicant:", err)
		return err
	}

	log.Info("New applicant inserted successfully:", newApplicant.ID)
	return nil
}
func GetApplicants() ([]models.Applicant,error){
 

    query := "SELECT id, nric, name, employment_status, sex, date_of_birth, household FROM applicants"
    rows, err := DB.Query(query)
	if err != nil {
		log.Error("Error executing query:", err)
		return nil, err
	}
	defer rows.Close() 

	var applicants []models.Applicant
	for rows.Next() {
		var applicant models.Applicant

		if err := rows.Scan(&applicant.ID,&applicant.NRIC, &applicant.Name, &applicant.EmploymentStatus, &applicant.Sex, &applicant.DateOfBirth, &applicant.Household); err != nil {
			log.Info("Error scanning row:", err)
			return nil, err
		}
		applicants = append(applicants, applicant) 
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Error("Error during row iteration:", err)
		return nil, err
	}

	return applicants, nil

}
func GetApplicantsMin() ([]models.Applicant,error){
 

    query := "SELECT id, nric, name, employment_status, sex, date_of_birth, household FROM applicants"
    rows, err := DB.Query(query)
	if err != nil {
		log.Error("Error executing query:", err)
		return nil, err
	}
	defer rows.Close() 

	var applicants []models.Applicant
	for rows.Next() {
		var applicant models.Applicant

		if err := rows.Scan(&applicant.ID,&applicant.NRIC, &applicant.Name, &applicant.EmploymentStatus, &applicant.Sex, &applicant.DateOfBirth, &applicant.Household); err != nil {
			log.Info("Error scanning row:", err)
			return nil, err
		}
		applicants = append(applicants, applicant) 
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Error("Error during row iteration:", err)
		return nil, err
	}

	return applicants, nil

}
func GetApplicantByID(applicant_id string ) (*models.Applicant,error){
 

    query := "SELECT id, nric, name, employment_status, sex, date_of_birth, household FROM applicants where id = ?"
    row := DB.QueryRow(query, applicant_id)

	var applicant models.Applicant

	err := row.Scan(&applicant.ID,&applicant.NRIC, &applicant.Name, &applicant.EmploymentStatus, &applicant.Sex, &applicant.DateOfBirth, &applicant.Household)
	if err != nil {
		log.Info("Error scanning row:", err)
		return nil, err
	}


	return &applicant, nil

}
func UpdateApplicant(id string, updatedApplicant models.Applicant) error {
	// Prepare the SQL query to update the applicant by ID
	query := `
		UPDATE applicants 
		SET nric =?, name = ?,  employment_status = ?, sex = ?, date_of_birth = ?, household = ? 
		WHERE id = ?`

	// Execute the query with the provided values
	_, err := DB.Exec(query, updatedApplicant.NRIC, updatedApplicant.Name, updatedApplicant.EmploymentStatus, 
		updatedApplicant.Sex, updatedApplicant.DateOfBirth, updatedApplicant.Household, id)
	
	if err != nil {
		log.Error("Error updating applicant:", err)
		return err
	}

	log.Info("Applicant updated successfully:", id)
	return nil
}
func DeleteApplicant(id string) error {
	// Prepare the SQL query to delete the applicant by ID
	query := `DELETE FROM applicants WHERE id = ?`

	// Execute the query
	_, err := DB.Exec(query, id)
	
	if err != nil {
		log.Error("Error deleting applicant:", err)
		return err
	}

	log.Info("Applicant deleted successfully:", id)
	return nil
}