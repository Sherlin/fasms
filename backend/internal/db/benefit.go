package db

import (
    "fasms/internal/models"
    log "github.com/sirupsen/logrus"

)
func CreateBenefit(newBenefit models.Benefit) error {
    query := `INSERT INTO benefits (id, name, amount, scheme_id) VALUES (?, ?, ?, ?)`
    _, err := DB.Exec(query, newBenefit.ID, newBenefit.Name, newBenefit.Amount, newBenefit.SchemeID)
    if err != nil {
        log.Error("Error inserting benefit:", err)
        return err
    }
    log.Info("New benefit inserted successfully:", newBenefit.ID)
    return nil
}

func GetBenefits() ([]models.Benefit, error) {
    query := `SELECT id, name, amount, scheme_id FROM benefits`
    rows, err := DB.Query(query)
    if err != nil {
        log.Error("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    var benefits []models.Benefit
    for rows.Next() {
        var benefit models.Benefit
        if err := rows.Scan(&benefit.ID, &benefit.Name, &benefit.Amount, &benefit.SchemeID); err != nil {
            log.Error("Error scanning row:", err)
            return nil, err
        }
        benefits = append(benefits, benefit)
    }
    return benefits, nil
}

func UpdateBenefit(id string, updatedBenefit models.Benefit) error {
    query := `UPDATE benefits SET name = ?, amount = ?, scheme_id = ? WHERE id = ?`
    _, err := DB.Exec(query, updatedBenefit.Name, updatedBenefit.Amount, updatedBenefit.SchemeID, id)
    if err != nil {
        log.Error("Error updating benefit:", err)
        return err
    }
    log.Info("Benefit updated successfully:", id)
    return nil
}

func DeleteBenefit(id string) error {
    query := `DELETE FROM benefits WHERE id = ?`
    _, err := DB.Exec(query, id)
    if err != nil {
        log.Error("Error deleting benefit:", err)
        return err
    }
    log.Info("Benefit deleted successfully:", id)
    return nil
}
