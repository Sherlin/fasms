package db

import (
    "fasms/internal/models"
    log "github.com/sirupsen/logrus"
)

func CreateScheme(newScheme models.Scheme) error {
    query := `INSERT INTO schemes (id, name, criteria, benefits) VALUES (?, ?, ?, ?)`
    _, err := DB.Exec(query, newScheme.ID, newScheme.Name, newScheme.Criteria, newScheme.Benefits)
    if err != nil {
        log.Error("Error inserting scheme:", err)
        return err
    }
    log.Info("New scheme inserted successfully:", newScheme.ID)
    return nil
}

func GetSchemes() ([]models.Scheme, error) {
    query := `SELECT id, name, criteria, benefits FROM schemes`
    rows, err := DB.Query(query)
    if err != nil {
        log.Error("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    var schemes []models.Scheme
    for rows.Next() {
        var scheme models.Scheme
        if err := rows.Scan(&scheme.ID, &scheme.Name, &scheme.Criteria, &scheme.Benefits); err != nil {
            log.Error("Error scanning row:", err)
            return nil, err
        }
        schemes = append(schemes, scheme)
    }
    return schemes, nil
}

func GetSchemesForApplicant() ([]models.Scheme, error) {
    query := `SELECT id, name, criteria, benefits FROM schemes`
    rows, err := DB.Query(query)
    if err != nil {
        log.Error("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()

    var schemes []models.Scheme
    for rows.Next() {
        var scheme models.Scheme
        if err := rows.Scan(&scheme.ID, &scheme.Name, &scheme.Criteria, &scheme.Benefits); err != nil {
            log.Error("Error scanning row:", err)
            return nil, err
        }
        schemes = append(schemes, scheme)
    }
    return schemes, nil
}

func UpdateScheme(id string, updatedScheme models.Scheme) error {
    query := `UPDATE schemes SET name = ?, criteria = ?, benefits = ? WHERE id = ?`
    _, err := DB.Exec(query, updatedScheme.Name, updatedScheme.Criteria, updatedScheme.Benefits, id)
    if err != nil {
        log.Error("Error updating scheme:", err)
        return err
    }
    log.Info("Scheme updated successfully:", id)
    return nil
}

func DeleteScheme(id string) error {
    query := `DELETE FROM schemes WHERE id = ?`
    _, err := DB.Exec(query, id)
    if err != nil {
        log.Error("Error deleting scheme:", err)
        return err
    }
    log.Info("Scheme deleted successfully:", id)
    return nil
}
