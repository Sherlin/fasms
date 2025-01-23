package db

import (
    "fasms/internal/models"
    log "github.com/sirupsen/logrus"

)


func CreateDependent(newDependent models.Dependent) error {
    query := `INSERT INTO dependents (id, name, employment_status, sex, relation, date_of_birth, parent_id) VALUES (?, ?, ?, ?, ?, ?, ?)`
    _, err := DB.Exec(query, newDependent.ID, newDependent.Name, newDependent.EmploymentStatus, newDependent.Sex, newDependent.Relation, newDependent.DateOfBirth, newDependent.ParentID)
    if err != nil {
        log.Error("Error inserting dependent:", err)
        return err
    }
    log.Info("New dependent inserted successfully:", newDependent.ID)
    return nil
}

func GetDependentsOfApplicant(parent_id string) ([]models.Dependent, error) {
    query := `SELECT id, name, employment_status, sex, relation, date_of_birth, parent_id FROM dependents where parent_id = ?`
    rows, err := DB.Query(query,parent_id)
    if err != nil {
        log.Error("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    var dependents []models.Dependent
    for rows.Next() {
        var dependent models.Dependent
        if err := rows.Scan(&dependent.ID, &dependent.Name, &dependent.EmploymentStatus, &dependent.Sex, &dependent.Relation, &dependent.DateOfBirth, &dependent.ParentID); err != nil {
            log.Error("Error scanning row:", err)
            return nil, err
        }
        dependents = append(dependents, dependent)
    }
    return dependents, nil
}

func UpdateDependent(id string, updatedDependent models.Dependent) error {
    query := `UPDATE dependents SET name = ?, employment_status = ?, sex = ?, relation = ?, date_of_birth = ?, parent_id = ? WHERE id = ?`
    _, err := DB.Exec(query, updatedDependent.Name, updatedDependent.EmploymentStatus, updatedDependent.Sex, updatedDependent.Relation, updatedDependent.DateOfBirth, updatedDependent.ParentID, id)
    if err != nil {
        log.Error("Error updating dependent:", err)
        return err
    }
    log.Info("Dependent updated successfully:", id)
    return nil
}

func DeleteDependent(id string) error {
    query := `DELETE FROM dependents WHERE id = ?`
    _, err := DB.Exec(query, id)
    if err != nil {
        log.Error("Error deleting dependent:", err)
        return err
    }
    log.Info("Dependent deleted successfully:", id)
    return nil
}