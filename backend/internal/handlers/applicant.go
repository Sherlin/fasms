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
func GetApplicantsMin(w http.ResponseWriter, r *http.Request) {

    applicants, err:= db.GetApplicantsMin()
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }    

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(applicants)
}
func GetApplicants(w http.ResponseWriter, r *http.Request) {

    applicants, err:= db.GetApplicants()
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }    
    for i:=0;i <len(applicants);i++{
        applicants[i].Dependents, err = db.GetDependentsOfApplicant(applicants[i].ID)
        if err != nil {
            http.Error(w, "Dependent of Applicant not found", http.StatusNotFound)
            return
        }    
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(applicants)
}
func GetApplicantByID(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    applicant, err:= db.GetApplicantByID(id)
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }    
    if (*applicant.Household != ""){
        applicant.Dependents, err = db.GetDependentsOfApplicant(applicant.ID)
        if err != nil {
            http.Error(w, "Dependent of Applicant not found", http.StatusNotFound)
            return
        } 
    }
    

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&applicant)
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

