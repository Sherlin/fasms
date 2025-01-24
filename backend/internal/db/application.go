package db

import (
    "fasms/internal/models"
    log "github.com/sirupsen/logrus"

)
func CreateApplication(newApplication models.Application) error {
    query := `INSERT INTO applications (id, date_of_application, scheme_id, applicant_id, status, disbursed, disbursed_date) VALUES (?, ?, ?, ?, ?, ?,?)`
    _, err := DB.Exec(query, newApplication.ID, newApplication.DateOfApplication, newApplication.SchemeID, newApplication.ApplicantID, newApplication.Status, newApplication.Disbursed, newApplication.DisbursedDate)
    if err != nil {
        log.Error("Error inserting application:", err)
        return err
    }
    log.Info("New application inserted successfully:", newApplication.ID)
    return nil
}

func GetApplications() ([]models.Application, error) {
    query := `SELECT id, date_of_application, scheme_id, applicant_id, status, disbursed, disbursed_date FROM applications`
    rows, err := DB.Query(query)
    if err != nil {
        log.Error("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    var applications []models.Application
    for rows.Next() {
        var application models.Application
        if err := rows.Scan(&application.ID, &application.DateOfApplication, &application.SchemeID, &application.ApplicantID, &application.Status, &application.Disbursed, &application.DisbursedDate); err != nil {
            log.Error("Error scanning row:", err)
            return nil, err
        }
        applications = append(applications, application)
    }
    return applications, nil
}

func UpdateApplication(id string, updatedApplication models.Application) error {
    query := `UPDATE applications SET scheme_id = ?, applicant_id = ?, status = ?, disbursed = ?, disbursed_date = ? WHERE id = ?`
    _, err := DB.Exec(query, updatedApplication.SchemeID, updatedApplication.ApplicantID, updatedApplication.Status, updatedApplication.Disbursed, updatedApplication.DisbursedDate, id)
    if err != nil {
        log.Error("Error updating application:", err)
        return err
    }
    log.Info("Application updated successfully:", id)
    return nil
}

func DeleteApplication(id string) error {
    query := `DELETE FROM applications WHERE id = ?`
    _, err := DB.Exec(query, id)
    if err != nil {
        log.Error("Error deleting application:", err)
        return err
    }
    log.Info("Application deleted successfully:", id)
    return nil
}