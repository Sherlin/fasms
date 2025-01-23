package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "fasms/internal/models"
    "fasms/internal/db"
)

func CreateScheme(w http.ResponseWriter, r *http.Request) {
    var scheme models.Scheme
    if err := json.NewDecoder(r.Body).Decode(&scheme); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    scheme.ID = uuid.New().String()
    if err := db.CreateScheme(scheme); err != nil {
        http.Error(w, "Insertion error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(scheme)
}

func GetSchemes(w http.ResponseWriter, r *http.Request) {
    schemes, err := db.GetSchemes()
    if err != nil {
        http.Error(w, "Schemes not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(schemes)
}
func GetSchemesForApplicant(w http.ResponseWriter, r *http.Request) {
    //id := mux.Vars(r)["applicant"]
    schemes, err := db.GetSchemes()
    if err != nil {
        http.Error(w, "Schemes not found", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(schemes)
}

func UpdateScheme(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updatedScheme models.Scheme
    if err := json.NewDecoder(r.Body).Decode(&updatedScheme); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := db.UpdateScheme(id, updatedScheme); err != nil {
        http.Error(w, "Scheme not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedScheme)
}

func DeleteScheme(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    if err := db.DeleteScheme(id); err != nil {
        http.Error(w, "Scheme not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
