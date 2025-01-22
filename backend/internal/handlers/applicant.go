package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)

var applicants = make(map[string]models.Applicant) // In-memory storage for demonstration

// CreateApplicant handles creating a new applicant.
func CreateApplicant(w http.ResponseWriter, r *http.Request) {
    var applicant models.Applicant
    if err := json.NewDecoder(r.Body).Decode(&applicant); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    applicant.ID = uuid.New().String() // Generate a unique ID
    err := db.CreateApplicant(applicant) 
    if err != nil {
        http.Error(w, "Insertion error", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(applicant)
}

// GetApplicant handles retrieving all applicants.
func GetApplicants(w http.ResponseWriter, r *http.Request) {

    applicants, err:= db.GetApplicants()
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }    

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(applicants)
}

// UpdateApplicant handles updating an existing applicant.
func UpdateApplicant(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedApplicant models.Applicant
    if err := json.NewDecoder(r.Body).Decode(&updatedApplicant); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    

    updatedApplicant.ID = id
    err := db.UpdateApplicant(id, updatedApplicant) 
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedApplicant)
}

// DeleteApplicant handles deleting an applicant by ID.
func DeleteApplicant(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    err := db.DeleteApplicant(id) 
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }
   
    w.WriteHeader(http.StatusNoContent)
}

// Helper function to format a date as dd-mm-yyyy
/*
func formatDate(date *time.Time) string {
    if date == nil {
        return ""
    }
    return date.Format("02-01-2006") // Format as dd-mm-yyyy
}*/