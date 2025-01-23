package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)

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

func GetApplications(w http.ResponseWriter, r *http.Request) {
    applications, err := db.GetApplications()
    if err != nil {
        http.Error(w, "Applications not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(applications)
}

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

func DeleteApplication(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := db.DeleteApplication(id); err != nil {
        http.Error(w, "Application not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}