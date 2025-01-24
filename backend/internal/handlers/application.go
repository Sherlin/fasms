package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)

// CreateApplication godoc
// @Summary Create a new application
// @Description Add a new application to the database
// @Tags applications
// @Accept json
// @Produce json
// @Param application body models.Application true "Application data"
// @Success 201 {object} models.Application
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Insertion error"
// @Router /api/applications [post]
func CreateApplication(w http.ResponseWriter, r *http.Request) {
    var application models.Application
    if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    application.ID = uuid.New().String()
    if err := db.CreateApplication(application); err != nil {
        http.Error(w, "Insertion error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(application)
}
// GetApplications godoc
// @Summary Get all applications
// @Description Retrieve all applications from the database
// @Tags applications
// @Produce json
// @Success 200 {array} models.Application
// @Failure 500 {string} string "Applications not found"
// @Router /api/applications [get]
func GetApplications(w http.ResponseWriter, r *http.Request) {
    applications, err := db.GetApplications()
    if err != nil {
        http.Error(w, "Applications not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(applications)
}

// UpdateApplication godoc
// @Summary Update an application
// @Description Update the details of an existing application
// @Tags applications
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Param application body models.Application true "Updated application data"
// @Success 200 {object} models.Application
// @Failure 400 {string} string "Invalid request payload"
// @Failure 404 {string} string "Application not found"
// @Router /api/applications/{id} [put]
func UpdateApplication(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedApplication models.Application
    if err := json.NewDecoder(r.Body).Decode(&updatedApplication); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := db.UpdateApplication(id, updatedApplication); err != nil {
        http.Error(w, "Application not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedApplication)
}

// DeleteApplication godoc
// @Summary Delete an application
// @Description Delete an application by its unique ID
// @Tags applications
// @Param id path string true "Application ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "Application not found"
// @Router /api/applications/{id} [delete]
func DeleteApplication(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := db.DeleteApplication(id); err != nil {
        http.Error(w, "Application not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}