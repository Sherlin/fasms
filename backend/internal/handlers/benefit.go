package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)
func CreateBenefit(w http.ResponseWriter, r *http.Request) {
    var benefit models.Benefit
    if err := json.NewDecoder(r.Body).Decode(&benefit); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    benefit.ID = uuid.New().String()
    if err := db.CreateBenefit(benefit); err != nil {
        http.Error(w, "Insertion error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(benefit)
}

func GetBenefits(w http.ResponseWriter, r *http.Request) {
    benefits, err := db.GetBenefits()
    if err != nil {
        http.Error(w, "Benefits not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(benefits)
}

func UpdateBenefit(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedBenefit models.Benefit
    if err := json.NewDecoder(r.Body).Decode(&updatedBenefit); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := db.UpdateBenefit(id, updatedBenefit); err != nil {
        http.Error(w, "Benefit not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedBenefit)
}

func DeleteBenefit(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := db.DeleteBenefit(id); err != nil {
        http.Error(w, "Benefit not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
