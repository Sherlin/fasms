package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)


func CreateDependent(w http.ResponseWriter, r *http.Request) {
    var dependent models.Dependent
    if err := json.NewDecoder(r.Body).Decode(&dependent); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    dependent.ID = uuid.New().String()
    if err := db.CreateDependent(dependent); err != nil {
        http.Error(w, "Insertion error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(dependent)
}
/*
func GetDependents(w http.ResponseWriter, r *http.Request) {
    dependents, err := db.GetDependents()
    if err != nil {
        http.Error(w, "Dependents not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(dependents)
}
*/
func UpdateDependent(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedDependent models.Dependent
    if err := json.NewDecoder(r.Body).Decode(&updatedDependent); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := db.UpdateDependent(id, updatedDependent); err != nil {
        http.Error(w, "Dependent not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedDependent)
}

func DeleteDependent(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := db.DeleteDependent(id); err != nil {
        http.Error(w, "Dependent not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
