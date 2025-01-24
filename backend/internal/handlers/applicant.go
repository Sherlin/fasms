package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)


// CreateApplicant godoc
// @Summary Create a new applicant
// @Description Add a new applicant to the database
// @Tags applicants
// @Accept json
// @Produce json
// @Param applicant body models.Applicant true "Applicant data"
// @Success 201 {object} models.Applicant
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Insertion error"
// @Router /api/applicants [post]
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

// GetApplicants godoc
// @Summary Get detailed list of applicants
// @Description Retrieve all applicants with their dependents
// @Tags applicants
// @Produce json
// @Success 200 {array} models.Applicant
// @Failure 404 {string} string "Applicants not found"
// @Router /api/applicants [get]
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

// GetApplicantByID godoc
// @Summary Get applicant by ID
// @Description Retrieve an applicant by their unique ID
// @Tags applicants
// @Produce json
// @Param id path string true "Applicant ID"
// @Success 200 {object} models.Applicant
// @Failure 404 {string} string "Applicant not found"
// @Router /api/applicants/{id} [get]
func GetApplicantByID(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    applicant, err:= db.GetApplicantByID(id)
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }    
    if ( applicant.Household != nil && *applicant.Household != ""){
        applicant.Dependents, err = db.GetDependentsOfApplicant(applicant.ID)
        if err != nil {
            http.Error(w, "Dependent of Applicant not found", http.StatusNotFound)
            return
        } 
    }
    

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&applicant)
}

// UpdateApplicant godoc
// @Summary Update an applicant
// @Description Update the details of an existing applicant
// @Tags applicants
// @Accept json
// @Produce json
// @Param id path string true "Applicant ID"
// @Param applicant body models.Applicant true "Updated applicant data"
// @Success 200 {object} models.Applicant
// @Failure 400 {string} string "Invalid request payload"
// @Failure 404 {string} string "Applicant not found"
// @Router /api/applicants/{id} [put]
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

// DeleteApplicant godoc
// @Summary Delete an applicant
// @Description Delete an applicant by their unique ID
// @Tags applicants
// @Param id path string true "Applicant ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "Applicant not found"
// @Router /api/applicants/{id} [delete]
func DeleteApplicant(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    err := db.DeleteApplicant(id) 
    if err != nil {
        http.Error(w, "Applicant not found", http.StatusNotFound)
        return
    }
   
    w.WriteHeader(http.StatusNoContent)
}

